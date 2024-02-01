package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "gokuk"
)

func GetPostgresConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("can not connect to db. err:", err)
	}

	return db
}

func ClosePostgresConnection(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("failed to get database object. err:", err)
		}

		err = sqlDB.Close()
		if err != nil {
			log.Fatal("failed to close database connection. err:", err)
		}
	}
}
