package application

type GetItemOrderRequest struct {
	ProductName  string `json:"productName"`
	ItemsOrdered int    `json:"items-ordered"`
}

type GetItemOrderResponse struct {
	ProductOrdered *OrderedItems `json:"product"`
	Err            string        `json:"err,omitempty"`
}

type PostProductRequest struct {
	Product *Product
}
