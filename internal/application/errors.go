package application

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	ProductNotFound       = "Product doesn't exist"
	InvalidProductRequest = "Invalid Request no Product supplied"
	InvalidOrderItem      = "Invalid Request no Item orders supplied, please supply value for 'ordered' parameter"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (re *RequestError) Error() string {
	return re.Err.Error()
}
func EncodeError(_ context.Context, err string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ProductNotFound:
		w.WriteHeader(http.StatusNotFound)
	case InvalidProductRequest:
		w.WriteHeader(http.StatusBadRequest)
	case InvalidOrderItem:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err,
	})
}
