package configs

import (
	"log"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "gokuk"
)

func GetPostgresConnection() *pgx.Conn {
	config := pgx.ConnConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: dbname,
	}
	conn, err := pgx.Connect(config)

	if err != nil {
		log.Fatal("can not connect to db. err:", err)
	}

	return conn
}

func ClosePostgresConnection(conn *pgx.Conn) {
	if conn.IsAlive() {
		defer conn.Close()
	}
}
