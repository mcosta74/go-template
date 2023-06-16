package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kit/kit/metrics"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/mcosta74/change-me/endpoints"
)

func MakeHTTPHandler(endpoints endpoints.Set, counter metrics.Counter) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerFinalizer(metricsFinalizer(counter)),
	}

	placeOrderHandler := kithttp.NewServer(
		endpoints.PlaceOrder,
		decodePlaceOrderRequest,
		kithttp.EncodeJSONResponse,
		opts...,
	)
	confirmOrderHandler := kithttp.NewServer(
		endpoints.ConfirmOrder,
		decodeConfirmOrderRequest,
		kithttp.EncodeJSONResponse,
		opts...,
	)
	cancelOrderHandler := kithttp.NewServer(
		endpoints.CancelOrder,
		decodeCancelOrderRequest,
		kithttp.EncodeJSONResponse,
		opts...,
	)
	getOrderHandler := kithttp.NewServer(
		endpoints.GetOrder,
		decodeGetOrderRequest,
		kithttp.EncodeJSONResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Method(http.MethodPost, "/orders", placeOrderHandler)
	r.Method(http.MethodGet, "/orders/{id}", getOrderHandler)
	r.Method(http.MethodPost, "/orders/{id}/confirm", confirmOrderHandler)
	r.Method(http.MethodPost, "/orders/{id}/cancel", cancelOrderHandler)

	return r
}
