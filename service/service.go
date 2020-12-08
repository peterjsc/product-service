package service

import (
	"context"
	"errors"
	"product-service/calculate"
	"product-service/internal/application"
	"product-service/internal/pkg/mongo/database"
	"product-service/repository"
	"strconv"
)

type service struct{}

var (
	ctx   = context.Background()
	db, _ = database.New(ctx)
)

type ProductService interface {
	ValidateProduct(product string) error
	ValidateOrder(orderAmount string) (int, error)
	GetProductOrdered(ctx context.Context, request application.GetItemOrderRequest) (*application.GetItemOrderResponse, error)
	PostProduct(ctx context.Context, request application.PostProductRequest) error
}

var (
	prodRepo = repository.NewProductRepository(ctx, db)
)

func NewProductService() ProductService {
	return &service{}
}

func (*service) ValidateProduct(product string) error {
	if product == "" {
		err := errors.New(application.InvalidProductRequest)
		return err
	}
	return nil
}

func (*service) ValidateOrder(orderAmount string) (int, error) {
	if orderAmount == "" {
		err := errors.New(application.InvalidOrderItem)
		return 0, err
	}

	itemsOrders, err := strconv.Atoi(orderAmount)
	if err != nil {
		err := errors.New(application.ErrConvertingInt)
		return 0, err
	}

	if itemsOrders == 0 {
		err := errors.New(application.InvalidOrderNum)
		return 0, err
	}

	return itemsOrders, nil
}

func (*service) GetProductOrdered(ctx context.Context, request application.GetItemOrderRequest) (*application.GetItemOrderResponse, error) {

	prodOrder, err := prodRepo.GetProduct(ctx, request.ProductName)
	if err != nil {
		return nil, err
	}

	orderedItems := calculate.CalcItemsWanted(request.ItemsOrdered, prodOrder.ItemPacks)

	orderInfo := &application.OrderedItems{
		ProductName: prodOrder.ProductName,
		ItemPacks:   orderedItems,
	}

	return &application.GetItemOrderResponse{ProductOrdered: orderInfo}, nil
}

func (*service) PostProduct(ctx context.Context, request application.PostProductRequest) error {

	prodOrder, _ := prodRepo.GetProduct(ctx, request.Product.ProductName)
	if prodOrder != nil {
		err := errors.New(application.ProductDuplicate)
		return err
	}

	addErr := prodRepo.CreateProduct(ctx, request.Product)
	if addErr != nil {
		return addErr
	}

	return nil
}
