package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				slog.Info(r.URL.Path, "time:", time.Since(start))
			}()

			f(w, r)
		}
	}
}
