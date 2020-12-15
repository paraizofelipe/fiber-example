package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/paraizofelipe/fiber-example/setting"
)

func ConnectDB() {
	var err error
	DB, err = sqlx.Connect(setting.Database.DBDriver, setting.Database.DBAddress)
	if err != nil {
		log.Panic(err)
	}
}
