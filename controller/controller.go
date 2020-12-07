package controller

import (
	"context"
	"encoding/json"
	"gymshark-interview/internal/application"
	"net/http"

	"gymshark-interview/service"

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
		ProductID:    product,
		ItemsOrdered: orderAmount,
	}

	newsArticles, err := psrv.GetProductOrdered(ctx, req)
	if err != nil {
		application.EncodeError(ctx, err.Error(), response)
		return
	}

	json.NewEncoder(response).Encode(newsArticles)
}
