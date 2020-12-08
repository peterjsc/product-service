package repository

import (
	"context"
	"errors"
	"fmt"
	"product-service/internal/application"
	"product-service/internal/pkg/mongo/database"

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
	GetProduct(ctx context.Context, productID string) (*application.Product, error)
	CreateProduct(ctx context.Context, product *application.Product) error
}

func NewProductRepository(ctx context.Context, db database.DBFramework) ProductRepository {
	collection := db.Database.Collection(ProductCollection)

	return &Repository{dbCol: collection}
}

func (pr *Repository) GetProduct(ctx context.Context, productID string) (*application.Product, error) {

	filter := bson.M{
		"product-name": productID,
	}

	var product *application.Product

	resErr := pr.dbCol.FindOne(context.TODO(), filter).Decode(&product)
	if resErr != nil {
		if resErr == mongo.ErrNoDocuments {
			return nil, errors.New(application.ProductNotFound)
		}
		return nil, resErr
	}

	return product, nil
}

func (pr *Repository) CreateProduct(ctx context.Context, product *application.Product) error {

	addProduct := bson.M{
		"product-name": product.ProductName,
		"item-packs":   product.ItemPacks,
	}

	productResult, err := pr.dbCol.InsertOne(ctx, addProduct)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted %v documents into episode collection!\n", productResult.InsertedID)
	return nil
}
