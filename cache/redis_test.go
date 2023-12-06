package cache

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redismock/v9"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisCache(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	cache := RedisCache{
		client: client,
		prefix: "test:",
		empty:  "empty",
	}

	// Test client initialization
	assert.NotNil(t, cache.client, "Redis client should not be nil")

	// Test prefix value
	assert.Equal(t, "test:", cache.prefix, "Prefix value should be 'test:'")

	// Test empty value
	assert.Equal(t, "empty", cache.empty, "Empty value should be 'empty'")
}

func TestRedisCache_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient, mock := redismock.NewClientMock()
	mockLogger := log.NewHelper(log.DefaultLogger)

	cache := RedisCache{
		client: mockClient,
		prefix: "test:",
		empty:  "empty",
		log:    mockLogger,
	}

	ctx := context.TODO()
	key := "test_key"
	value := "test_value"
	mock.ExpectGet(cache.prefix + key).SetVal(value)
	getValue, _ := cache.Get(ctx, key)
	assert.Equal(t, value, getValue, "Get value should be equal to 'test_value'")
}

func TestRedisCache_MGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient, mock := redismock.NewClientMock()
	mockLogger := log.NewHelper(log.DefaultLogger)

	cache := RedisCache{
		client: mockClient,
		prefix: "test:",
		empty:  "empty",
		log:    mockLogger,
	}

	ctx := context.TODO()
	keys := []string{"test_key1", "test_key2"}
	values := []interface{}{"test_value1", "test_value2"}
	mock.ExpectMGet(cache.handleKeys(keys...)...).SetVal(values)
	getValues, _ := cache.MGet(ctx, keys...)
	mapValues := make(map[string]string, len(keys))
	for i, key := range keys {
		mapValues[cache.prefix+key], _ = values[i].(string)
	}
	assert.Equal(t, mapValues, getValues, "Get values should be equal to 'test_value1' and 'test_value2'")
}
