package domain

type Payment struct {
	PaymentId   string `json:"payment_id"`
	CustomerId  string `json:"customer_id"`
	Amount      string `json:"amount"`
	PaymentType string `json:"payment_type"`
}
