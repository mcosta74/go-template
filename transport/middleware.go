package transport

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kit/kit/metrics"
)

func instrumentingMdw(counter metrics.Counter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wr := middleware.NewWrapResponseWriter(w, 1)

			defer func() {
				counter.With("status_code", fmt.Sprint(wr.Status())).Add(1)
			}()

			next.ServeHTTP(wr, r)
		})
	}
}
