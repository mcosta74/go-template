package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mcosta74/change-me/repository"
	"golang.org/x/exp/slog"
)

type Order repository.Order
type OrderStatus repository.OrderStatus

var (
	StatusPending   = OrderStatus(repository.StatusPending)
	StatusConfirmed = OrderStatus(repository.StatusConfirmed)
	StatusCanceled  = OrderStatus(repository.StatusCanceled)
)

type OrderSvc interface {
	GetOrder(ctx context.Context, id int) (*Order, error)
	PlaceOrder(ctx context.Context, o *Order) (*Order, error)
	ConfirmOrder(ctx context.Context, id int) (*Order, error)
	CancelOrder(ctx context.Context, id int) (*Order, error)
}

func New(r repository.OrderRepository, logger *slog.Logger) OrderSvc {
	var svc OrderSvc
	{
		svc = &orderSvc{
			r: r,
		}
		svc = newLoggingMdw(logger, svc)
	}
	return svc
}

type orderSvc struct {
	r repository.OrderRepository
}

func (s *orderSvc) GetOrder(ctx context.Context, id int) (*Order, error) {
	v, err := s.r.GetOrder(ctx, id)
	return (*Order)(v), err
}

func (s *orderSvc) PlaceOrder(ctx context.Context, o *Order) (*Order, error) {
	now := time.Now()

	o.Created = now
	o.LastUpdate = now
	o.Status = repository.StatusPending

	v, err := s.r.CreateOrder(ctx, (*repository.Order)(o))
	return (*Order)(v), err
}

func (s *orderSvc) ConfirmOrder(ctx context.Context, id int) (*Order, error) {
	o, err := s.updateOrderStatus(ctx, id, StatusConfirmed)
	if err != nil {
		return nil, fmt.Errorf("confirm order failed: %w", err)
	}
	return o, nil
}

func (s *orderSvc) CancelOrder(ctx context.Context, id int) (*Order, error) {
	o, err := s.updateOrderStatus(ctx, id, StatusCanceled)
	if err != nil {
		return nil, fmt.Errorf("confirm order failed: %w", err)
	}
	return o, nil
}

func (s *orderSvc) updateOrderStatus(ctx context.Context, id int, newStatus OrderStatus) (*Order, error) {
	o, err := s.r.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	o.Status = repository.OrderStatus(newStatus)
	o.LastUpdate = time.Now()

	v, err := s.r.UpdateStatus(ctx, o)
	return (*Order)(v), err
}
