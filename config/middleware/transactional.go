package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionalDbKeyType struct{}

func Transactional(pool *pgxpool.Pool) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tx, err := pool.Begin(r.Context())
			if err != nil {
				slog.Error("unable to start transaction", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			defer func() {
				if p := recover(); p != nil {
					tx.Rollback(r.Context())
					slog.Error("transaction rolled back", "error", p)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					panic(p)
				} else if err != nil {
					tx.Rollback(r.Context())
					slog.Error("transaction rolled back", "error", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				} else {
					err = tx.Commit(r.Context())
					if err != nil {
						slog.Error("failed to commit transaction", "error", err)
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					}
				}
			}()

			// add transaction to the context
			ctx := context.WithValue(r.Context(), &TransactionalDbKeyType{}, tx)
			r = r.WithContext(ctx)

			f(w, r)
		}
	}
}
