package main

import (
	"fmt"
	"log"

	app "github.com/zhora-ip/libraries-management-system/intenal/app/http_app"
	sqldb "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"
	"github.com/zhora-ip/libraries-management-system/pkg"
)

var (
	pathDB = "configs/database.yaml"
)

func main() {

	db := &sqldb.Config{}

	pkg.ParseConfig(db, pathDB)

	databaseURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host,
		db.Port,
		db.Username,
		db.Password,
		db.DBName,
	)

	cfg := &app.Config{
		DatabaseURL: databaseURL,
	}
	
	if err := app.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
