package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/exp/slog"
)

func loggingMdw(logger *slog.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (response any, err error) {
			defer func(dt time.Time) {
				logger.Debug("endpoint called", "success", err == nil, "took", time.Since(dt))
			}(time.Now())

			return next(ctx, request)
		}
	}
}
