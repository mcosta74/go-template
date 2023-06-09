package repository

import (
	"context"
	"fmt"
	"sync"
)

func NewInMemRepository() Repository {
	return &inMemRepo{
		items: make(map[int]*Item),
	}
}

type inMemRepo struct {
	mutex  sync.RWMutex
	lastID int
	items  map[int]*Item
}

func (r *inMemRepo) List(ctx context.Context, filters map[string][]string) ([]*Item, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	items := make([]*Item, 0, len(r.items))
	for _, item := range r.items {
		items = append(items, item)
	}
	return items, nil
}

func (r *inMemRepo) Insert(ctx context.Context, item *Item) (*Item, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, rItem := range r.items {
		if rItem.Code == item.Code {
			return nil, fmt.Errorf("%w: code already exists", ErrConflict)
		}
	}
	r.lastID++
	item.ID = r.lastID
	r.items[item.ID] = item
	return item, nil
}

func (r *inMemRepo) Get(ctx context.Context, id int) (*Item, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if item, ok := r.items[id]; ok {
		return item, nil
	}
	return nil, ErrNotFound
}

func (r *inMemRepo) Update(ctx context.Context, item *Item) (*Item, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	rItem, ok := r.items[item.ID]
	if !ok {
		return nil, ErrNotFound
	}

	rItem.Description = item.Description
	return rItem, nil
}

func (r *inMemRepo) Delete(ctx context.Context, id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.items[id]; ok {
		delete(r.items, id)
		return nil
	}
	return ErrNotFound
}
