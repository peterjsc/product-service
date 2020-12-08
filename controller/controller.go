package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"product-service/internal/application"

	"product-service/service"

	"github.com/gorilla/mux"
)

type controller struct{}

var (
	ctx = context.Background()
)

var (
	psrv = service.NewProductService()
)

type ProductController interface {
	GetItemsOrdered(response http.ResponseWriter, request *http.Request)
	PostProduct(response http.ResponseWriter, request *http.Request)
}

func NewProductController() ProductController {
	return &controller{}
}

func (*controller) GetItemsOrdered(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	queries := request.URL.Query()

	v := mux.Vars(request)

	product := v["product"]

	itemsOrdered := queries.Get("ordered")

	orderAmount, err := psrv.ValidateOrder(itemsOrdered)
	if err != nil {
		application.EncodeError(ctx, err.Error(), response)
		return
	}

	valErr := psrv.ValidateProduct(product)
	if valErr != nil {
		application.EncodeError(ctx, valErr.Error(), response)
		return
	}

	req := application.GetItemOrderRequest{
		ProductName:  product,
		ItemsOrdered: orderAmount,
	}

	newsArticles, err := psrv.GetProductOrdered(ctx, req)
	if err != nil {
		application.EncodeError(ctx, err.Error(), response)
		return
	}

	json.NewEncoder(response).Encode(newsArticles)
}

func (*controller) PostProduct(response http.ResponseWriter, request *http.Request) {

	if request.Header.Get("Content-Type") == "" {
		msg := "Content-Type header is not application/json"
		application.EncodeError(ctx, msg, response)
		return
	}

	v := mux.Vars(request)

	product := &application.Product{
		ProductName: v["product"],
	}

	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		application.EncodeError(ctx, application.InvalidRequestBody, response)
		return
	}

	if product.ItemPacks == nil {
		application.EncodeError(ctx, application.InvalidRequestBody, response)
		return
	}

	prodRequest := application.PostProductRequest{
		Product: product,
	}

	respErr := psrv.PostProduct(ctx, prodRequest)
	if respErr != nil {
		application.EncodeError(ctx, respErr.Error(), response)
		return
	}

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.WriteHeader(http.StatusCreated)

	fmt.Fprintf(response, "Product is now created")
}
