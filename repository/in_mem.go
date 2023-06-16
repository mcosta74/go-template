package repository

import (
	"context"
	"sync"
)

func NewInMemRepository() OrderRepository {
	return &inMemRepository{
		orders: make(map[int]*Order),
	}
}

type inMemRepository struct {
	idSequence int
	m          sync.RWMutex
	orders     map[int]*Order
}

func (r *inMemRepository) GetOrder(ctx context.Context, id int) (*Order, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if o, ok := r.orders[id]; ok {
		return o, nil
	}
	return nil, ErrNotFound
}

func (r *inMemRepository) CreateOrder(ctx context.Context, o *Order) (*Order, error) {
	r.m.Lock()
	defer r.m.Unlock()

	r.idSequence++
	o.ID = r.idSequence

	r.orders[o.ID] = o
	return o, nil
}

func (r *inMemRepository) UpdateStatus(ctx context.Context, o *Order) (*Order, error) {
	r.m.Lock()
	defer r.m.Unlock()

	order, ok := r.orders[o.ID]
	if !ok {
		return nil, ErrNotFound
	}
	order.LastUpdate = o.LastUpdate
	order.Status = o.Status
	return order, nil
}
