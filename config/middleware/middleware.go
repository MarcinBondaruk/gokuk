package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Pipe(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}

	return handler
}
