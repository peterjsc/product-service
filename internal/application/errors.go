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
	InvalidOrderNum       = "Please supply an order amount greater than 0"
	ErrConvertingInt      = "Error converting string to int"
	ContentTypeError      = "Content-Type header is not application/json"
	InvalidRequestBody    = "Please supply a body in json for request or in correct format"
	ProductDuplicate      = "Product already exists"
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
	case InvalidProductRequest, InvalidOrderItem, InvalidOrderNum, ErrConvertingInt, InvalidRequestBody:
		w.WriteHeader(http.StatusBadRequest)
	case ContentTypeError:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	case ProductDuplicate:
		w.WriteHeader(http.StatusConflict)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err,
	})
}
