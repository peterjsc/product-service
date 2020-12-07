package application

type GetItemOrderRequest struct {
	ProductID    string `json:"productID"`
	ItemsOrdered int    `json:"items-ordered"`
}

type GetItemOrderResponse struct {
	ProductOrdered *OrderedItems `json:"product"`
	Err            string        `json:"err,omitempty"`
}
