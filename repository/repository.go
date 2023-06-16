package repository

import (
	"context"
	"errors"
	"time"
)

type OrderStatus int

const (
	StatusPending OrderStatus = iota
	StatusConfirmed
	StatusCanceled
)

type Order struct {
	ID          int         `json:"id,omitempty"`
	Description string      `json:"description,omitempty"`
	Created     time.Time   `json:"created,omitempty"`
	Status      OrderStatus `json:"status,omitempty"`
	LastUpdate  time.Time   `json:"last_update,omitempty"`
}

var (
	ErrNotFound = errors.New("not found")
)

type OrderRepository interface {
	GetOrder(ctx context.Context, id int) (*Order, error)
	CreateOrder(ctx context.Context, o *Order) (*Order, error)
	UpdateStatus(ctx context.Context, o *Order) (*Order, error)
}
