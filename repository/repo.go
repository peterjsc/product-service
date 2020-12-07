package repository

import (
	"context"
	"errors"
	"gymshark-interview/calculate"
	"gymshark-interview/internal/application"
	"gymshark-interview/internal/pkg/mongo/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	dbCol *mongo.Collection
}

const (
	ProductCollection = "product"
)

type ProductRepository interface {
	GetProductOrdered(ctx context.Context, productID string, itemsOrdered int) (*application.OrderedItems, error)
}

func NewProductRepository(ctx context.Context, db database.DBFramework) ProductRepository {
	collection := db.Database.Collection(ProductCollection)

	return &Repository{dbCol: collection}
}

func (pr *Repository) GetProductOrdered(ctx context.Context, productID string, itemsOrdered int) (*application.OrderedItems, error) {

	filter := bson.M{
		"product-name": productID,
	}

	var product *application.Product

	resErr := pr.dbCol.FindOne(context.TODO(), filter).Decode(&product)
	if resErr != nil {
		if resErr == mongo.ErrNoDocuments {
			return &application.OrderedItems{}, errors.New(application.ProductNotFound)
		}
		return &application.OrderedItems{}, resErr
	}

	orderedItems := calculate.CalcItemsWanted(itemsOrdered, product.ItemPacks)

	orderInfo := &application.OrderedItems{
		ProductName: product.ProductName,
		ItemPacks:   orderedItems,
	}

	return orderInfo, nil
}
