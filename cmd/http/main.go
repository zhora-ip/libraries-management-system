package main

import (
	"fmt"
	"log"

	kafkaservice "github.com/zhora-ip/libraries-management-system/infrastructure/kafka"
	app "github.com/zhora-ip/libraries-management-system/intenal/app/http_app"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/redis"
	sqldb "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"
	"github.com/zhora-ip/libraries-management-system/pkg"
)

var (
	pathDB    = "configs/database.yaml"
	pathKafka = "configs/kafka.yaml"
	pathRedis = "configs/redis.yaml"
)

func main() {

	var (
		db       = &sqldb.Config{}
		kafkaCfg = &kafkaservice.Config{}
		redisCfg = &redis.Config{}
		cfg      = &app.Config{}
	)

	pkg.ParseConfig(db, pathDB)
	databaseURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host,
		db.Port,
		db.Username,
		db.Password,
		db.DBName,
	)

	pkg.ParseConfig(kafkaCfg, pathKafka)
	pkg.ParseConfig(redisCfg, pathRedis)

	cfg.DatabaseURL = databaseURL
	cfg.KafkaCfg = kafkaCfg
	cfg.RedisCfg = redisCfg

	if err := app.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
