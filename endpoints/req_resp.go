package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/mcosta74/change-me/service"
)

type ListItemsRequests struct {
	Filters map[string][]string `json:"filters,omitempty"`
}

type ListItemsResponse struct {
	V []*service.Item
}

func (r ListItemsResponse) MarshalJSON() ([]byte, error) {
	temp := struct {
		Count int             `json:"count"`
		V     []*service.Item `json:"v"`
	}{Count: len(r.V), V: r.V}

	return json.Marshal(temp)
}

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
	Item *service.Item `json:"item"`
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
