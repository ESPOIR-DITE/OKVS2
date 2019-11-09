package order

type Orders struct {
	Id         string `json:"id"`
	CustomerId string `json:"customer_id"`
	Date       string `json:"date"`
}

type OrderLine struct {
	Linecode    string `json:"linecode"`
	ItemNumber  string `json:"item_number"`
	OrderNumber string `json:"order_number"`
}
