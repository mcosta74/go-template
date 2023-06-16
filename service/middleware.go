package service

import (
	"context"

	"golang.org/x/exp/slog"
)

type loggingMdw struct {
	logger *slog.Logger
	next   OrderSvc
}

func newLoggingMdw(logger *slog.Logger, next OrderSvc) OrderSvc {
	return &loggingMdw{
		logger: logger,
		next:   next,
	}
}

func (mdw *loggingMdw) GetOrder(ctx context.Context, id int) (o *Order, err error) {
	defer func() {
		mdw.logger.Debug("Get Order", "id", id, "success", err != nil)
	}()
	o, err = mdw.next.GetOrder(ctx, id)
	return
}

func (mdw *loggingMdw) PlaceOrder(ctx context.Context, o *Order) (newOrder *Order, err error) {
	defer func() {
		mdw.logger.Debug("Place Order", "description", o.Description, "success", err != nil)
	}()
	newOrder, err = mdw.next.PlaceOrder(ctx, o)
	return
}

func (mdw *loggingMdw) ConfirmOrder(ctx context.Context, id int) (o *Order, err error) {
	defer func() {
		mdw.logger.Debug("Confirm Order", "oder_id", id, "success", err != nil)
	}()
	o, err = mdw.next.ConfirmOrder(ctx, id)
	return
}

func (mdw *loggingMdw) CancelOrder(ctx context.Context, id int) (o *Order, err error) {
	defer func() {
		mdw.logger.Debug("Cancel Order", "oder_id", id, "success", err != nil)
	}()
	o, err = mdw.next.CancelOrder(ctx, id)
	return
}
