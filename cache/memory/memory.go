package memory

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var (
	defaultExpiration      string = "5m"
	defaultCleanupInterval string = "10m"
)

var (
	// ErrKeyNotFound 缓存键未找到
	ErrKeyNotFound = errors.New("cache Key doesn't found")
)

type memoryCache struct {
	Expiration      time.Duration
	CleanupInterval time.Duration
	gcache          *gocache.Cache
}

func New(opt map[string]interface{}) *memoryCache {
	if exp, ok := opt["DefaultExpiration"]; ok {
		defaultExpiration = exp.(string)
	}
	if intv, ok := opt["CleanupInterval"]; ok {
		defaultCleanupInterval = intv.(string)
	}
	durExp, err := time.ParseDuration(defaultExpiration)
	if err != nil {
		panic("memcache expiration setting error")
	}
	durIntv, err := time.ParseDuration(defaultCleanupInterval)
	if err != nil {
		panic("memcache cleanup interval setting error")
	}
	return &memoryCache{
		gcache:          gocache.New(durExp, durIntv),
		Expiration:      durExp,
		CleanupInterval: durIntv,
	}
}

func (c *memoryCache) Get(key string) (interface{}, error) {
	if val, found := c.gcache.Get(key); found {
		return val, nil
	}
	return nil, ErrKeyNotFound
}
func (c *memoryCache) Read(key string, outPtr interface{}) error {
	curr := reflect.ValueOf(outPtr)
	if curr.Kind() != reflect.Ptr {
		return errors.New("the value must be pointer type")
	}
	val, err := c.Get(key)
	if err != nil {
		outPtr = nil
		return err
	}
	oldVal := reflect.ValueOf(val)
	if oldVal.Kind() == reflect.Ptr {
		curr.Elem().Set(oldVal.Elem())
	} else {
		curr.Elem().Set(oldVal)
	}
	return nil
}
func (c *memoryCache) Set(key string, value interface{}, secondsLifetime int64) error {
	if secondsLifetime > 0 {
		secs, err := time.ParseDuration(fmt.Sprintf("%vs", secondsLifetime))
		if err != nil {
			secs = c.Expiration
		}
		c.gcache.Set(key, value, secs)
		return nil
	}
	if secondsLifetime == 0 {
		c.gcache.Set(key, value, c.Expiration)
		return nil
	}
	c.gcache.Set(key, value, gocache.NoExpiration)
	return nil
}
func (c *memoryCache) Delete(key string) error {
	c.gcache.Delete(key)
	return nil
}
