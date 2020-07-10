package accounting

import "time"

type Accounting struct {
	ItemId   string  `json:"itemId"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Payment struct {
	PaymentNumber string `json:"payment_number"`
	PaymentTypeId string `json:"payment_type_id"`
	Amount        string `json:"amount"`
	OrderNumber   string `json:"order_number"`
}
type AccountChange struct {
	Id              string    `json:"id"`
	ItemId          string    `json:"itemId"`
	Date            time.Time `json:"date"`
	InitialQuantity int       `json:"initialQuantity"`
	PostQuantity    int       `json:"postQuantity"`
	Description     string    `json:"description"`
}
