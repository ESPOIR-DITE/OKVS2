package domain

type Order struct {
	OrderId  string `json:"order_id"`
	Customer string `json:"customer"`
	Items    []Item `json:"items"`
	Date     string `json:"date"`
}
