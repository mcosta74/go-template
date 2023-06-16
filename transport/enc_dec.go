package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/mcosta74/change-me/endpoints"
)

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	if coder, ok := err.(kithttp.StatusCoder); ok {
		code = coder.StatusCode()
	}

	temp := struct {
		Detail string `json:"detail,omitempty"`
	}{Detail: err.Error()}

	contentType := "application/json"
	body, err := json.Marshal(temp)
	if err != nil {
		body = []byte(err.Error())
		contentType = "application/text"
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func getIntUrlParam(r *http.Request, key string) (int, error) {
	param := chi.URLParam(r, key)
	val, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("%q is not valid value for parameter %q", param, key)
	}
	return val, nil
}

func decodePlaceOrderRequest(ctx context.Context, r *http.Request) (any, error) {
	var req endpoints.PlaceOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, endpoints.NewDecodeError(err)
	}
	return req, nil
}

func decodeConfirmOrderRequest(ctx context.Context, r *http.Request) (any, error) {
	val, err := getIntUrlParam(r, "id")
	if err != nil {
		return nil, endpoints.NewDecodeError(err)
	}
	return endpoints.ConfirmOrderRequest{ID: val}, nil
}

func decodeCancelOrderRequest(ctx context.Context, r *http.Request) (any, error) {
	val, err := getIntUrlParam(r, "id")
	if err != nil {
		return nil, endpoints.NewDecodeError(err)
	}
	return endpoints.CancelOrderRequest{ID: val}, nil
}

func decodeGetOrderRequest(ctx context.Context, r *http.Request) (any, error) {
	val, err := getIntUrlParam(r, "id")
	if err != nil {
		return nil, endpoints.NewDecodeError(err)
	}
	return endpoints.GetOrderRequest{ID: val}, nil
}
