package service

import (
	"context"

	"github.com/mcosta74/change-me/repository"
)

type Item struct {
	ID          int    `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

func itemFromRepository(item *repository.Item) *Item {
	return &Item{
		ID:          item.ID,
		Code:        item.Code,
		Description: item.Description,
	}
}

func itemToRepository(item *Item) *repository.Item {
	return &repository.Item{
		ID:          item.ID,
		Code:        item.Code,
		Description: item.Description,
	}
}

type Service interface {
	ListItems(ctx context.Context, filters map[string][]string) ([]*Item, error)
	CreateItem(ctx context.Context, item *Item) (*Item, error)
	ReadItem(ctx context.Context, id int) (*Item, error)
	UpdateItem(ctx context.Context, item *Item) (*Item, error)
	DeleteItem(ctx context.Context, id int) error
}

func New(repo repository.Repository) Service {
	return &service{repo: repo}
}

type service struct {
	repo repository.Repository
}

func (svc *service) ListItems(ctx context.Context, filters map[string][]string) ([]*Item, error) {
	v, err := svc.repo.List(ctx, filters)
	if err != nil {
		return []*Item{}, err
	}

	items := make([]*Item, 0, len(v))
	for _, item := range v {
		items = append(items, itemFromRepository(item))
	}
	return items, nil
}

func (svc *service) CreateItem(ctx context.Context, item *Item) (*Item, error) {
	v, err := svc.repo.Insert(ctx, itemToRepository(item))
	if err != nil {
		return nil, err
	}
	return itemFromRepository(v), nil
}

func (svc *service) ReadItem(ctx context.Context, id int) (*Item, error) {
	v, err := svc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return itemFromRepository(v), nil
}

func (svc *service) UpdateItem(ctx context.Context, item *Item) (*Item, error) {
	v, err := svc.repo.Update(ctx, itemToRepository(item))
	if err != nil {
		return nil, err
	}
	return itemFromRepository(v), nil
}

func (svc *service) DeleteItem(ctx context.Context, id int) error {
	return svc.repo.Delete(ctx, id)
}
