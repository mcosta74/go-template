package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/mcosta74/change-me/endpoints"
)

func MakeHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createHandler := kithttp.NewServer(
		endpoints.CreateItem,
		decodeCreateItem,
		kithttp.EncodeJSONResponse,
		opts...,
	)
	readHandler := kithttp.NewServer(
		endpoints.ReadItem,
		decodeReadItem,
		kithttp.EncodeJSONResponse,
		opts...,
	)
	updateHandler := kithttp.NewServer(
		endpoints.UpdateItem,
		decodeUpdateItem,
		kithttp.EncodeJSONResponse,
		opts...,
	)
	deleteHandler := kithttp.NewServer(
		endpoints.DeleteItem,
		decodeDeleteItem,
		kithttp.EncodeJSONResponse,
		opts...,
	)

	r.Method(http.MethodPost, "/items", createHandler)
	r.Method(http.MethodGet, "/items/{id}", readHandler)
	r.Method(http.MethodPut, "/items/{id}", updateHandler)
	r.Method(http.MethodDelete, "/items/{id}", deleteHandler)

	return r
}