package database

import (
	"context"
	"fmt"
	"log/slog"

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

func InitConnPool() (*pgxpool.Pool, error) {
	connString := formatConnectionString()
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		slog.Error("can not parse db connection string", "error:", err)
		return nil, err
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		uuid.Register(conn.TypeMap())

		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		slog.Error("db failed to connect", "error:", err)
	}

	return pool, nil
}
