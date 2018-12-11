package redis

type Config struct {
	Addr     string // 缓存地址
	Password string // 密码
	Database string
	Prefix   string // 键值前缀
}
