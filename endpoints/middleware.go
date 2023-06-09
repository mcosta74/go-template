package endpoints

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"github.com/mcosta74/change-me/repository"
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

func instrumentingMdw(duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (response any, err error) {
			defer func(dt time.Time) {
				duration.With("success", fmt.Sprint(err == nil)).Observe(time.Since(dt).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func errorsMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		resp, err := next(ctx, request)
		if err != nil {
			switch {
			case errors.Is(err, repository.ErrNotFound):
				err = NewNotFoundError(err)

			case errors.Is(err, repository.ErrConflict):
				err = NewConflictError(err)
			}
		}
		return resp, err
	}
}
