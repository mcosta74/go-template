package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"github.com/mcosta74/change-me/service"
	"golang.org/x/exp/slog"
)

type Set struct {
	PlaceOrder   endpoint.Endpoint
	ConfirmOrder endpoint.Endpoint
	CancelOrder  endpoint.Endpoint
	GetOrder     endpoint.Endpoint
}

func New(svc service.OrderSvc, logger *slog.Logger, duration metrics.Histogram) Set {
	var placeOrder endpoint.Endpoint
	{
		placeOrder = makePlaceOrderEndpoint(svc)
		placeOrder = instrumentingMdw(duration.With("PlaceOrder"))(placeOrder)
		placeOrder = loggingMdw(logger.With("name", "PlaceOrder"))(placeOrder)
		placeOrder = errorsMiddleware(placeOrder)
	}

	var confirmOrder endpoint.Endpoint
	{
		confirmOrder = makeConfirmOrderEndpoint(svc)
		confirmOrder = instrumentingMdw(duration.With("ConfirmOrder"))(confirmOrder)
		confirmOrder = loggingMdw(logger.With("name", "ConfirmOrder"))(confirmOrder)
		confirmOrder = errorsMiddleware(confirmOrder)
	}

	var cancelOrder endpoint.Endpoint
	{
		cancelOrder = makeCancelOrderEndpoint(svc)
		cancelOrder = instrumentingMdw(duration.With("CancelOrder"))(cancelOrder)
		cancelOrder = loggingMdw(logger.With("name", "CancelOrder"))(cancelOrder)
		cancelOrder = errorsMiddleware(cancelOrder)
	}

	var getOrder endpoint.Endpoint
	{
		getOrder = makeGetOrderEndpoint(svc)
		getOrder = instrumentingMdw(duration.With("GetOrder"))(getOrder)
		getOrder = loggingMdw(logger.With("name", "GetOrder"))(getOrder)
		getOrder = errorsMiddleware(getOrder)
	}

	return Set{
		PlaceOrder:   placeOrder,
		ConfirmOrder: confirmOrder,
		CancelOrder:  cancelOrder,
		GetOrder:     getOrder,
	}
}

func makePlaceOrderEndpoint(svc service.OrderSvc) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(PlaceOrderRequest)

		v, err := svc.PlaceOrder(ctx, &service.Order{Description: req.Description})
		if err != nil {
			return nil, err
		}
		return &PlaceOrderResponse{V: v}, nil
	}
}

func makeConfirmOrderEndpoint(svc service.OrderSvc) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(ConfirmOrderRequest)

		v, err := svc.ConfirmOrder(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return &ConfirmOrderResponse{V: v}, nil
	}
}

func makeCancelOrderEndpoint(svc service.OrderSvc) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(CancelOrderRequest)

		v, err := svc.CancelOrder(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return &CancelOrderResponse{V: v}, nil
	}
}

func makeGetOrderEndpoint(svc service.OrderSvc) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(GetOrderRequest)

		v, err := svc.GetOrder(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return &GetOrderResponse{V: v}, nil
	}
}
