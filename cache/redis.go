// Package cache RedisCache
package cache

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/yimoka/api/fault"
)

// RedisCache  实现 Cache 接口 RedisCache
type RedisCache struct {
	// redis 连接池
	client *redis.Client
	log    *log.Helper
	// 缓存前缀
	prefix string
	// 空值
	empty string
}

// NewRedisCache 创建 RedisCache
func NewRedisCache(client *redis.Client, prefix string, logger log.Logger) *RedisCache {
	if client == nil {
		panic("redis client 不能为空")
	}
	return &RedisCache{
		client: client,
		prefix: prefix,
		log:    log.NewHelper(logger),
	}
}

// IsEmpty 判断缓存是否为空
func (r *RedisCache) IsEmpty(value string) bool {
	return value == r.empty
}

// Get 获取缓存
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, r.prefix+key).Result()
	if err != nil {
		if err == redis.Nil {
			return val, fault.ErrorNotFound("缓存不存在")
		}
		r.log.Errorf("redis get key: %s error: %v", key, err)
	}
	return val, err
}

// MGet 批量获取缓存
func (r *RedisCache) MGet(ctx context.Context, keys ...string) (map[string]string, error) {
	if len(keys) == 0 {
		return nil, fault.ErrorBadRequest("参数错误")
	}
	return r.mGet(ctx, r.handleKeys(keys...)...)
}

// PrefixGet 前置匹配获取
func (r *RedisCache) PrefixGet(ctx context.Context, prefix string, scanCount int64) (map[string]string, error) {
	var cursor uint64
	var keys []string
	var err error
	data := make(map[string]string)
	for {
		keys, cursor, err = r.client.Scan(ctx, cursor, r.prefix+prefix+"*", scanCount).Result()
		if err != nil {
			r.log.Errorf("redis scan prefix: %s error: %v", prefix, err)
			return nil, fault.ErrorInternalServerError("前置匹配获取缓存失败")
		}
		if len(keys) > 0 {
			values, err := r.mGet(ctx, keys...)
			if err != nil {
				return nil, err
			}
			data = lo.Assign(data, values)
		}
		if cursor == 0 {
			break
		}
	}
	return data, nil
}

// Set 设置缓存
func (r *RedisCache) Set(ctx context.Context, key string, val string, expiration time.Duration) error {
	err := r.client.Set(ctx, r.prefix+key, val, expiration).Err()
	if err != nil {
		r.log.Errorf("redis set key: %s error: %v", key, err)
		return fault.ErrorInternalServerError("设置缓存失败")
	}
	return nil
}

// MSet 批量设置缓存
func (r *RedisCache) MSet(ctx context.Context, data map[string]string, expiration time.Duration) error {
	length := len(data)
	if length == 0 {
		return fault.ErrorBadRequest("参数错误")
	}
	// 使用 pipeline 批量设置
	pipe := r.client.Pipeline()
	defer pipe.Discard()
	for key, value := range data {
		pipe.Set(ctx, r.prefix+key, value, expiration)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis pipe.Set data: %v error: %v", data, err)
		return fault.ErrorInternalServerError("批量设置缓存失败")
	}
	return nil
}

// SetEmpty 设置空值 防止缓存穿透
func (r *RedisCache) SetEmpty(ctx context.Context, key string, expiration time.Duration) error {
	return r.Set(ctx, r.prefix+key, r.empty, expiration)
}

// MSetEmpty 批量设置空值
func (r *RedisCache) MSetEmpty(ctx context.Context, keys []string, expiration time.Duration) error {
	if len(keys) == 0 {
		return fault.ErrorBadRequest("参数错误")
	}
	// 使用 pipeline 批量设置
	pipe := r.client.Pipeline()
	defer pipe.Discard()
	for _, key := range keys {
		pipe.Set(ctx, r.prefix+key, r.empty, expiration)
	}
	_, err := pipe.Exec(ctx)
	return err
}

// Del 删除缓存
func (r *RedisCache) Del(ctx context.Context, key string) error {
	err := r.client.Del(ctx, r.prefix+key).Err()
	if err != nil {
		r.log.Errorf("redis del key: %s error: %v", key, err)
		return fault.ErrorInternalServerError("删除缓存失败")
	}
	return nil
}

// MDel 批量删除缓存
func (r *RedisCache) MDel(ctx context.Context, keys ...string) error {
	return r.mDel(ctx, r.handleKeys(keys...)...)
}

// PrefixDel 前缀
func (r *RedisCache) PrefixDel(ctx context.Context, prefix string, scanCount int64) error {
	var cursor uint64
	var keys []string
	var err error
	for {
		keys, cursor, err = r.client.Scan(ctx, cursor, r.prefix+prefix+"*", scanCount).Result()
		if err != nil {
			r.log.Errorf("redis scan prefix: %s error: %v", prefix, err)
			return fault.ErrorInternalServerError("前置匹配删除缓存失败")
		}
		if len(keys) > 0 {
			err = r.mDel(ctx, keys...)
			if err != nil {
				return err
			}
		}
		if cursor == 0 {
			break
		}
	}
	return nil
}

// Clear 清空缓存
func (r *RedisCache) Clear(ctx context.Context, scanCount int64) error {
	var cursor uint64
	var keys []string
	var err error
	for {
		keys, cursor, err = r.client.Scan(ctx, cursor, r.prefix+"*", scanCount).Result()
		if err != nil {
			r.log.Errorf("redis scan error: %v", err)
			return fault.ErrorInternalServerError("清空缓存失败")
		}
		if len(keys) > 0 {
			err = r.mDel(ctx, keys...)
			if err != nil {
				return err
			}
		}
		if cursor == 0 {
			break
		}
	}
	return nil
}

// Close 关闭缓存
func (r *RedisCache) Close() error {
	return r.client.Close()
}

// GetType 获取缓存类型
func (r *RedisCache) GetType() string {
	return "redis"
}

// 处理 keys 前缀
func (r *RedisCache) handleKeys(keys ...string) []string {
	if r.prefix == "" {
		return keys
	}
	if len(keys) == 0 {
		return nil
	}
	newKeys := make([]string, len(keys))
	for i, key := range keys {
		newKeys[i] = r.prefix + key
	}
	return newKeys
}

// mGet 批量获取缓存 不处理前缀
func (r *RedisCache) mGet(ctx context.Context, keys ...string) (map[string]string, error) {
	length := len(keys)
	if length == 0 {
		return nil, fault.ErrorBadRequest("参数错误")
	}
	if length <= 1000 {
		values, err := r.client.MGet(ctx, keys...).Result()
		if err != nil {
			r.log.Errorf("redis mGet keys: %v error: %v", keys, err)
			return nil, fault.ErrorInternalServerError("批量获取缓存失败")
		}
		strMap := map[string]string{}
		for i, value := range values {
			if value != nil {
				if v, ok := value.(string); ok {
					strMap[keys[i]] = v
				}
			}
		}
		return strMap, nil
	}
	// 分批获取
	data := map[string]string{}
	for i := 0; i < length; i += 1000 {
		end := i + 1000
		if end > length {
			end = length
		}
		values, err := r.client.MGet(ctx, keys[i:end]...).Result()
		if err != nil {
			r.log.Errorf("redis mGet keys: %v error: %v", keys, err)
			return nil, fault.ErrorInternalServerError("批量获取缓存失败")
		}
		for j, value := range values {
			if value != nil {
				if v, ok := value.(string); ok {
					data[keys[i+j]] = v
				}
			}
		}
	}
	return data, nil
}

// mDel 删除不处理前缀
func (r *RedisCache) mDel(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return fault.ErrorBadRequest("参数错误")
	}
	err := r.client.Del(ctx, keys...).Err()
	if err != nil {
		r.log.Errorf("redis mDel keys: %v error: %v", keys, err)
		return fault.ErrorInternalServerError("批量删除缓存失败")
	}
	return nil
}

// type Data struct {
// 	cache Cache
// }

// func NewData(cache Cache) *Data {
// 	return &Data{
// 		cache: cache,
// 	}
// }

// var redisData = NewData(&RedisCache{})
