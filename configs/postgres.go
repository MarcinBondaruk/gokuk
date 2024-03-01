package configs

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "gokuk"
)

func formatConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func GetPostgresConnection() *pgx.Conn {
	config := formatConnectionString()
	// change conntor to v5 version
	conn, err := pgx.Connect(context.Background(), config)

	if err != nil {
		log.Fatal("can not connect to db. err:", err)
	}

	return conn
}

func ClosePostgresConnection(conn *pgx.Conn) {
	if !conn.IsClosed() {
		defer conn.Close(context.Background())
	}
}
