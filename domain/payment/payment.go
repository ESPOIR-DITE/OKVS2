package payment

type Payment struct {
	PaymentNumber string `json:"payment_number"`
	PaymentTypeId string `json:"payment_type_id"`
	Amount        string `json:"amount"`
	OrderNumber   string `json:"order_number"`
}
