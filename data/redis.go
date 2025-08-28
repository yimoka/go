// Package data redis.go
package data

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/yimoka/go/config"
)

// GetRedisClient returns a Redis client based on the provided configuration.
func GetRedisClient(config *config.Redis) redis.UniversalClient {
	var rdb redis.UniversalClient
	if config.IsSingle {
		rdb = redis.NewClient(&redis.Options{
			Addr:         config.Addr,
			Password:     config.Password,
			DB:           int(config.Db),
			DialTimeout:  config.DialTimeout.AsDuration(),
			WriteTimeout: config.WriteTimeout.AsDuration(),
			ReadTimeout:  config.ReadTimeout.AsDuration(),
		})
	} else {
		rdb = getRedisClusterClient(config)
	}
	if config.IsTrace {
		if err := redisotel.InstrumentTracing(rdb); err != nil {
			panic(err)
		}
	}
	return rdb
}

// GetRedisClient returns a Redis client based on the provided configuration.
func getRedisClusterClient(config *config.Redis) *redis.ClusterClient {
	addrs := config.Addrs
	if len(addrs) == 0 && config.Addr != "" {
		addrs = []string{config.Addr}
	}
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        addrs,
		Password:     config.Password,
		DialTimeout:  config.DialTimeout.AsDuration(),
		WriteTimeout: config.WriteTimeout.AsDuration(),
		ReadTimeout:  config.ReadTimeout.AsDuration(),
	})
	if config.IsTrace {
		if err := redisotel.InstrumentTracing(rdb); err != nil {
			panic(err)
		}
	}
	return rdb
}
