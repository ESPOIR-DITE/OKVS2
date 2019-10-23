package order

type Orders struct {
	OrderNumber string `json:"order_number"`
	Customer_id string `json:"customer_id"`
	Date        string `json:"date"`
}

type OrderLine struct {
	Linecode    string `json:"linecode"`
	ItemNumber  string `json:"item_number"`
	OrderNumber string `json:"order_number"`
}
