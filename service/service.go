package service

import (
	"context"
	"errors"
	"fmt"
	"gymshark-interview/internal/application"
	"gymshark-interview/internal/pkg/mongo/database"
	"gymshark-interview/repository"
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

	prodOrder, err := prodRepo.GetProductOrdered(ctx, request.ProductName, request.ItemsOrdered)
	if err != nil {
		return nil, err
	}
	return &application.GetItemOrderResponse{ProductOrdered: prodOrder}, nil
}

func (*service) PostProduct(ctx context.Context, request application.PostProductRequest) error {

	fmt.Println(request.Product.ProductName)
	fmt.Println(request.Product.ItemPacks)

	err := prodRepo.CreateProduct(ctx, request.Product)
	if err != nil {
		return err
	}

	return nil
}
