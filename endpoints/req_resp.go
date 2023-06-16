package endpoints

import (
	"net/http"

	"github.com/mcosta74/change-me/service"
)

type GetOrderRequest struct {
	ID int `json:"id"`
}

type GetOrderResponse struct {
	V *service.Order `json:"v,omitempty"`
}

type PlaceOrderRequest struct {
	Description string `json:"description,omitempty"`
}

type PlaceOrderResponse struct {
	V *service.Order `json:"v,omitempty"`
}

func (r *PlaceOrderResponse) StatusCode() int {
	return http.StatusCreated
}

type ConfirmOrderRequest struct {
	ID int `json:"id"`
}

type ConfirmOrderResponse struct {
	V *service.Order `json:"v,omitempty"`
}

type CancelOrderRequest struct {
	ID int `json:"id"`
}

type CancelOrderResponse struct {
	V *service.Order `json:"v,omitempty"`
}
