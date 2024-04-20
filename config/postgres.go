package config

import (
	"context"
	"fmt"
	"log"

	"github.com/MarcinBondaruk/gokuk/database/extensions/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO: use envs
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

func GetPostgresConnection() *pgxpool.Pool {
	connString := formatConnectionString()
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal("can not parse db connection string. err:", err)
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		uuid.Register(conn.TypeMap())

		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		log.Fatal("can not connect to db. err:", err)
	}

	log.Print("db initialized")
	return pool
}

func ClosePostgresConnection(pool *pgxpool.Pool) {
	pool.Close()
}
