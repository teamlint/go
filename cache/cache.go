package cache

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
