package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Pipe(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		h = m(h)
	}

	return h
}
