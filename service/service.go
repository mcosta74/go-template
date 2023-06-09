package service

import "context"

type Item struct {
	ID          int    `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type ItemService interface {
	CreateItem(ctx context.Context, item *Item) (*Item, error)
	ReadItem(ctx context.Context, id int) (*Item, error)
	UpdateItem(ctx context.Context, item *Item) (*Item, error)
	DeleteItem(ctx context.Context, id int) error
}

func NewItemService() ItemService {
	return &itemSvc{}
}

type itemSvc struct{}

func (svc *itemSvc) CreateItem(ctx context.Context, item *Item) (*Item, error) {
	return &Item{}, nil
}

func (svc *itemSvc) ReadItem(ctx context.Context, id int) (*Item, error) {
	return &Item{}, nil
}

func (svc *itemSvc) UpdateItem(ctx context.Context, item *Item) (*Item, error) {
	return &Item{}, nil
}

func (svc *itemSvc) DeleteItem(ctx context.Context, id int) error {
	return nil
}
