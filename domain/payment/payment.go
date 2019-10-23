package payment

type Payment struct {
	PaymentNUmber string `json:"payment_n_umber"`
	PaymentType   string `json:"payment_type"`
	Amount        string `json:"amount"`
	OrderNumber   string `json:"order_number"`
}
