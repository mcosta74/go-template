package transport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/metrics"
	kithttp "github.com/go-kit/kit/transport/http"
)

func metricsFinalizer(counter metrics.Counter) kithttp.ServerFinalizerFunc {
	return func(ctx context.Context, code int, r *http.Request) {
		counter.With("status_code", fmt.Sprint(code)).Add(1)
	}
}
