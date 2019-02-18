package redis

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

var (
	// ErrRedisClosed redis 已关闭
	ErrRedisClosed = errors.New("redis is already closed")
	// ErrKeyNotFound 缓存键未找到
	ErrKeyNotFound           = errors.New("cache Key doesn't found")
	defaultExpiration string = "5m"
)

// redisCache 缓存实现
type redisCache struct {
	Config     *Config
	pool       *redigo.Pool
	Connected  bool
	Expiration time.Duration
}

// New 创建redisCache实例
func New(opt map[string]interface{}) *redisCache {
	setting := Config{}
	if addr, ok := opt["Addr"]; ok {
		setting.Addr = addr.(string)
	}
	if pwd, ok := opt["Password"]; ok {
		setting.Password = pwd.(string)
	}
	if database, ok := opt["Database"]; ok {
		setting.Database = database.(string)
	}
	if prefix, ok := opt["Prefix"]; ok {
		setting.Prefix = prefix.(string)
	}
	if exp, ok := opt["DefaultExpiration"]; ok {
		defaultExpiration = exp.(string)
	}
	durExp, err := time.ParseDuration(defaultExpiration)
	if err != nil {
		panic("memcache expiration setting error")
	}
	r := &redisCache{pool: &redigo.Pool{}, Expiration: durExp, Config: &setting}
	r.connect()
	_, err = r.PingPong()
	if err != nil {
		log.Fatalf("error connecting to redis: %v", err)
		return nil
	}
	runtime.SetFinalizer(r, close)
	return r
}

func (r *redisCache) connect() {
	c := r.Config

	pool := &redigo.Pool{IdleTimeout: 240 * time.Second, MaxIdle: 3}
	// 空闲连接心跳检查
	pool.TestOnBorrow = func(c redigo.Conn, t time.Time) error {
		_, err := c.Do("PING")
		return err
	}
	// 连接
	if c.Database != "" {
		pool.Dial = func() (redigo.Conn, error) {
			red, err := dial("tcp", c.Addr, c.Password)
			if err != nil {
				return nil, err
			}
			if _, err = red.Do("SELECT", c.Database); err != nil {
				red.Close()
				return nil, err
			}
			return red, err
		}
	} else {
		pool.Dial = func() (redigo.Conn, error) {
			return dial("tcp", c.Addr, c.Password)
		}
	}
	r.Connected = true
	r.pool = pool
}

func dial(network string, addr string, pass string) (redigo.Conn, error) {
	c, err := redigo.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	if pass != "" {
		if _, err = c.Do("AUTH", pass); err != nil {
			c.Close()
			return nil, err
		}
	}
	return c, err
}

func close(r *redisCache) error {
	if r.pool != nil {
		return r.pool.Close()
	}
	return ErrRedisClosed
}

func (r *redisCache) PingPong() (bool, error) {
	c := r.pool.Get()
	defer c.Close()
	msg, err := c.Do("PING")
	if err != nil || msg == nil {
		return false, err
	}
	return (msg == "PONG"), nil
}

// Get 获取指定缓存键的值
func (r *redisCache) Get(key string) (value interface{}, err error) {
	data, e := r.get(key)
	if e != nil {
		err = e
		return
	}

	if e := DefaultSerializer.Unmarshal(data.([]byte), &value); e != nil {
		log.Fatalf("unable to unmarshal value of key: '%s': %v", key, e)
		err = e
		return
	}
	return
}
func (r *redisCache) get(key string) (interface{}, error) {
	c := r.pool.Get()
	defer c.Close()
	if err := c.Err(); err != nil {
		return nil, err
	}

	redisVal, err := c.Do("GET", r.Config.Prefix+key)

	if err != nil {
		return nil, err
	}
	if redisVal == nil {
		return nil, ErrKeyNotFound
	}
	return redisVal, nil
}

// Read 读取指定键到变量
func (r *redisCache) Read(key string, outPtr interface{}) error {
	if reflect.ValueOf(outPtr).Kind() != reflect.Ptr {
		return errors.New("the value must be pointer type")
	}
	data, err := r.get(key)
	if err != nil {
		return err
	}

	if err = DefaultSerializer.Unmarshal(data.([]byte), outPtr); err != nil {
		log.Fatalf("unable to unmarshal value of key: '%s': %v", key, err)
		return fmt.Errorf("unable to unmarshal value of key: '%s': %v", key, err)
	}
	return nil
}

// Set 设置值
func (r *redisCache) Set(key string, value interface{}, secondsLifetime int64) (err error) {
	valueBytes, e := DefaultSerializer.Marshal(value)
	if e != nil {
		log.Fatal(e)
		return e
	}
	return r.set(key, valueBytes, secondsLifetime)

}
func (r *redisCache) set(key string, valueBytes interface{}, secondsLifetime int64) (err error) {
	c := r.pool.Get()
	defer c.Close()
	if c.Err() != nil {
		return c.Err()
	}

	// if has expiration, then use the "EX" to delete the key automatically.
	if secondsLifetime > 0 {
		_, err = c.Do("SETEX", r.Config.Prefix+key, secondsLifetime, valueBytes)
		return
	}
	if secondsLifetime == 0 {
		_, err = c.Do("SETEX", r.Config.Prefix+key, int64(r.Expiration.Seconds()), valueBytes)
		return
	}
	_, err = c.Do("SET", r.Config.Prefix+key, valueBytes)
	return
}

// Delete 删除键
func (r *redisCache) Delete(key string) error {
	c := r.pool.Get()
	defer c.Close()
	if c.Err() != nil {
		return c.Err()
	}

	_, err := c.Do("DEL", r.Config.Prefix+key)
	return err
}
