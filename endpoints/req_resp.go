package endpoints

import (
	"net/http"

	"github.com/mcosta74/change-me/service"
)

type CreateItemRequest struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateItemResponse struct {
	V *service.Item `json:"v,omitempty"`
}

func (r CreateItemResponse) StatusCode() int {
	return http.StatusCreated
}

type ReadItemRequest struct {
	ID int `json:"id,omitempty"`
}

type ReadItemResponse struct {
	V *service.Item `json:"v,omitempty"`
}

type UpdateItemRequest struct {
	Item *service.Item `json:"item,omitempty"`
}

type UpdateItemResponse struct {
	V *service.Item `json:"v,omitempty"`
}

type DeleteItemRequest struct {
	ID int `json:"id,omitempty"`
}

type DeleteItemResponse struct {
}

func (r DeleteItemResponse) StatusCode() int {
	return http.StatusNoContent
}
