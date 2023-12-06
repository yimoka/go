// Package cache table
// 封装数据库表缓存获取与设置
package cache

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
	"github.com/yimoka/api/fault"
	"golang.org/x/sync/singleflight"
)

// Table 定义表格缓存结构体，用于缓存表格数据，减少数据库查询次数。
type Table struct {
	log          *log.Helper
	cache        Cache
	singleFlight *singleflight.Group
	expireTime   time.Duration
}

// TableContent 定义表格内容的接口
type TableContent[T comparable] struct {
	IsPreventPenetration bool
	GetKey               func(T) string
	GetValue             func(T) (string, error)
	GetValueMap          func(...T) (map[T]string, error)
}

// NewTable 初始化表格缓存
func NewTable(cache Cache, singleFlight *singleflight.Group, expireTimeSecond int32, logger *log.Helper) *Table {
	return &Table{
		cache:        cache,
		log:          logger,
		singleFlight: singleFlight,
		expireTime:   time.Duration(expireTimeSecond) * time.Second,
	}
}

// GetTableCache 获取获取
func GetTableCache[T comparable](ctx context.Context, t *Table, content *TableContent[T], param T) (string, error) {
	cacheKey := content.GetKey(param)
	cacheValue, err := t.cache.Get(ctx, cacheKey)
	if err != nil {
		if !fault.IsNotFound(err) {
			return "", err
		}
		dbVal, dbErr, _ := t.singleFlight.Do(cacheKey, func() (interface{}, error) {
			detail, dErr := content.GetValue(param)
			if dErr != nil {
				if fault.IsNotFound(dErr) && content.IsPreventPenetration {
					// 防止缓存穿透
					sErr := t.cache.SetEmpty(ctx, cacheKey, t.expireTime)
					if sErr != nil {
						t.log.Errorf("setEmpty cache error: %v", sErr)
					}
				}
				return nil, dErr
			}
			sErr := t.cache.Set(ctx, cacheKey, detail, t.expireTime)
			if sErr != nil {
				t.log.Errorf("set cache error: %v", sErr)
			}
			return detail, nil
		})
		if dbErr != nil {
			return "", dbErr
		}
		cacheValue, _ = dbVal.(string)
	}
	if t.cache.IsEmpty(cacheValue) {
		return "", fault.ErrorNotFound("找不到或已删除，请检查您的参数")
	}
	return cacheValue, nil
}

// SetTableCache 设置缓存
func SetTableCache[T comparable](ctx context.Context, t *Table, content *TableContent[T], param T) error {
	cacheKey := content.GetKey(param)
	cacheValue, err := content.GetValue(param)
	if err != nil {
		if !fault.IsNotFound(err) && content.IsPreventPenetration {
			// 防止缓存穿透
			_ = t.cache.SetEmpty(ctx, cacheKey, t.expireTime)
		}
		return err
	}
	return t.cache.Set(ctx, cacheKey, cacheValue, t.expireTime)
}

// MGetTableCache 获取多个缓存
func MGetTableCache[T comparable](ctx context.Context, t *Table, content *TableContent[T], params ...T) (map[T]string, error) {
	keyToParam := make(map[string]T)
	cacheKeys := lo.Map(params, func(item T, index int) string {
		key := content.GetKey(item)
		keyToParam[key] = item
		return key
	})
	cacheValues, err := t.cache.MGet(ctx, cacheKeys...)
	if err != nil {
		return nil, err
	}

	// 去掉空值, cacheKey 转为原来的 param
	values := lo.MapKeys(lo.OmitBy(cacheValues, func(key string, value string) bool { return t.cache.IsEmpty(value) }),
		func(_ string, key string) T { return keyToParam[key] },
	)

	if len(cacheValues) == len(params) {
		return values, nil
	}

	// 从数据库中获取
	// 找到缓存中没有的 key
	var dbParams []T
	for _, cacheKey := range cacheKeys {
		if _, ok := cacheValues[cacheKey]; !ok {
			dbParams = append(dbParams, keyToParam[cacheKey])
		}
	}
	dbMap, gErr := content.GetValueMap(dbParams...)
	if gErr != nil {
		return nil, gErr
	}
	// 如果防止缓存穿透，找到参数有但数据库中没有的，就设置为空
	if content.IsPreventPenetration && len(dbParams) != len(dbMap) {
		emptyKeys := []string{}
		for _, dbParam := range dbParams {
			if _, ok := dbMap[dbParam]; !ok {
				emptyKeys = append(emptyKeys, content.GetKey(dbParam))
			}
		}
		if len(emptyKeys) > 0 {
			_ = t.cache.MSetEmpty(ctx, emptyKeys, t.expireTime)
		}
	}

	// 设置缓存
	if len(dbMap) == 0 {
		return values, nil
	}
	_ = mSetTableCache(ctx, t, content, dbMap)
	// 合并值
	return lo.Assign(values, dbMap), err
}

// MSetTableCache 设置多个缓存
func MSetTableCache[T comparable](ctx context.Context, t *Table, content *TableContent[T], params ...T) error {
	cacheValues, err := content.GetValueMap(params...)
	if err != nil {
		return err
	}
	return mSetTableCache(ctx, t, content, cacheValues)
}

// DelTableCache 删除缓存
func DelTableCache[T comparable](ctx context.Context, t *Table, content *TableContent[T], params ...T) error {
	cacheKeys := lo.Map(params, func(item T, index int) string {
		return content.GetKey(item)
	})
	return t.cache.MDel(ctx, cacheKeys...)
}

func mSetTableCache[T comparable](ctx context.Context, t *Table, content *TableContent[T], values map[T]string) error {
	cacheValues := lo.MapKeys(values, func(value string, key T) string {
		return content.GetKey(key)
	})
	return t.cache.MSet(ctx, cacheValues, t.expireTime)
}
