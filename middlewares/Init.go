package middlewares

import "net/http"

type Adapter func(http.Handler) http.Handler

func Adapt(handler http.Handler, adapters ...Adapter) http.Handler {
	for i := range adapters {
		handler = adapters[len(adapters)-1-i](handler)
	}
	return handler
}
