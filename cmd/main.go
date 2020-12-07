package main

import (
	"gymshark-interview/controller"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	defaultHTTPPort string = ":8080"
)

func main() {

	httpRouter := mux.NewRouter()
	prodController := controller.NewProductController()

	httpRouter.HandleFunc("/products/{product}", prodController.GetItemsOrdered).Methods("GET")

	http.ListenAndServe(defaultHTTPPort, httpRouter)
}
