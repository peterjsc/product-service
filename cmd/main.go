package main

import (
	"net/http"
	"product-service/controller"

	"github.com/gorilla/mux"
)

const (
	defaultHTTPPort string = ":8080"
)

func main() {

	httpRouter := mux.NewRouter()
	prodController := controller.NewProductController()

	httpRouter.HandleFunc("/products/{product}", prodController.GetItemsOrdered).Methods("GET")
	httpRouter.HandleFunc("/products/{product}", prodController.PostProduct).Methods("POST")

	http.ListenAndServe(defaultHTTPPort, httpRouter)
}
