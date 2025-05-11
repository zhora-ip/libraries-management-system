package app

import (
	kafkaservice "github.com/zhora-ip/libraries-management-system/infrastructure/kafka"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/redis"
)

type Config struct {
	DatabaseURL string
	KafkaCfg    *kafkaservice.Config
	RedisCfg    *redis.Config
}
