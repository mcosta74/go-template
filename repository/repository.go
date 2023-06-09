package repository

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
)

type Item struct {
	ID          int
	Code        string
	Description string
}

type Repository interface {
	List(ctx context.Context, filters map[string][]string) ([]*Item, error)
	Insert(ctx context.Context, item *Item) (*Item, error)
	Get(ctx context.Context, id int) (*Item, error)
	Update(ctx context.Context, item *Item) (*Item, error)
	Delete(ctx context.Context, id int) error
}
