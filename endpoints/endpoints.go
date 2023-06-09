package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"github.com/mcosta74/change-me/service"
	"golang.org/x/exp/slog"
)

type Endpoints struct {
	ListItems  endpoint.Endpoint
	CreateItem endpoint.Endpoint
	ReadItem   endpoint.Endpoint
	UpdateItem endpoint.Endpoint
	DeleteItem endpoint.Endpoint
}

func MakeEndpoints(svc service.ItemService, logger *slog.Logger, duration metrics.Histogram) Endpoints {
	var listItems endpoint.Endpoint
	{
		listItems = makeListItems(svc)
		listItems = instrumentingMdw(duration.With("name", "ListItems"))(listItems)
		listItems = loggingMdw(logger.With("name", "ListItems"))(listItems)
	}

	var createItem endpoint.Endpoint
	{
		createItem = makeCreateItem(svc)
		createItem = instrumentingMdw(duration.With("name", "CreateItem"))(createItem)
		createItem = loggingMdw(logger.With("name", "CreateItem"))(createItem)
	}

	var readItem endpoint.Endpoint
	{
		readItem = makeReadItem(svc)
		readItem = instrumentingMdw(duration.With("name", "ReadItem"))(readItem)
		readItem = loggingMdw(logger.With("name", "ReadItem"))(readItem)
	}

	var updateItem endpoint.Endpoint
	{
		updateItem = makeUpdateItem(svc)
		updateItem = instrumentingMdw(duration.With("name", "UpdateItem"))(updateItem)
		updateItem = loggingMdw(logger.With("name", "UpdateItem"))(updateItem)
	}

	var deleteItem endpoint.Endpoint
	{
		deleteItem = makeDeleteItem(svc)
		deleteItem = instrumentingMdw(duration.With("name", "DeleteItem"))(deleteItem)
		deleteItem = loggingMdw(logger.With("name", "DeleteItem"))(deleteItem)
	}

	return Endpoints{
		ListItems:  listItems,
		CreateItem: createItem,
		ReadItem:   readItem,
		UpdateItem: updateItem,
		DeleteItem: deleteItem,
	}
}

func makeListItems(svc service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(ListItemsRequests)

		v, err := svc.ListItems(ctx, req.Filters)
		if err != nil {
			return nil, err
		}
		return ListItemsResponse{V: v}, nil
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
