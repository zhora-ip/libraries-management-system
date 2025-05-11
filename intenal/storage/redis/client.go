package redis

import (
	"github.com/redis/go-redis/v9"
)

var (
	expirationTime = 1
)

type Config struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	TTL      int    `yaml:"ttl"`
}

type Redis struct {
	client *redis.Client
}

func New(ops *Config) *Redis {
	expirationTime = ops.TTL
	return &Redis{
		redis.NewClient(
			&redis.Options{
				Addr:     ops.Addr,
				Password: ops.Password,
				DB:       ops.DB,
			},
		),
	}
}
