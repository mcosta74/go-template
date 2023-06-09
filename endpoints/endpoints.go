package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mcosta74/change-me/service"
)

type Endpoints struct {
	CreateItem endpoint.Endpoint
	ReadItem   endpoint.Endpoint
	UpdateItem endpoint.Endpoint
	DeleteItem endpoint.Endpoint
}

func MakeEndpoints(svc service.ItemService) Endpoints {
	return Endpoints{
		CreateItem: makeCreateItem(svc),
		ReadItem:   makeReadItem(svc),
		UpdateItem: makeUpdateItem(svc),
		DeleteItem: makeDeleteItem(svc),
	}
}

func makeCreateItem(svc service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(CreateItemRequest)

		v, err := svc.CreateItem(ctx, &service.Item{Code: req.Code, Description: req.Description})
		if err != nil {
			return nil, err
		}
		return CreateItemResponse{v}, nil
	}
}

func makeReadItem(svc service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(ReadItemRequest)

		v, err := svc.ReadItem(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return ReadItemResponse{v}, nil
	}
}

func makeUpdateItem(svc service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(UpdateItemRequest)

		v, err := svc.UpdateItem(ctx, req.Item)
		if err != nil {
			return nil, err
		}
		return UpdateItemResponse{v}, nil
	}
}

func makeDeleteItem(svc service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(DeleteItemRequest)

		err := svc.DeleteItem(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return DeleteItemResponse{}, nil
	}
}
