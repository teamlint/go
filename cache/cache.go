package cache

import (
	"fmt"

	"github.com/teamlint/gox/cache/memory"
	"github.com/teamlint/gox/cache/redis"
)

const (
	// NoExpiration 没有有效期限制
	NoExpiration int64 = -1
	// 默认有效期
	DefaultExpiration int64 = 0
)

// Cache 缓存接口
type Cache interface {
	Get(key string) (interface{}, error)
	Read(key string, outPtr interface{}) error
	Set(key string, value interface{}, secondsLifetime int64) error
	Delete(key string) error
}

// Type 缓存类型
type Type int

const (
	TypeMemory Type = 1 // 本地内存缓存
	TypeRedis  Type = 2 // Redis缓存
)

// New 创建指定类型缓存实例
func New(t Type, opt ...map[string]interface{}) Cache {
	option := make(map[string]interface{})
	switch t {
	case TypeRedis:
		fmt.Println("redis cache")
		option = map[string]interface{}{
			"Addr":              "127.0.0.1:6379",
			"Password":          "",
			"Database":          "",
			"Prefix":            "",
			"DefaultExpiration": "20m",
		}
		if len(opt) > 0 {
			option = opt[0]
		}
		return redis.New(option)
	default:
		fmt.Println("memory cache")
		option = map[string]interface{}{
			"DefaultExpiration": "20m",
			"CleanupInterval":   "30m",
		}
		if len(opt) > 0 {
			option = opt[0]
		}
		return memory.New(option)
	}
}
