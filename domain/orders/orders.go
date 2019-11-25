package orders

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
type orderHelper struct {
	ItemNumber string `json:"itemNumber"`
	Quantity   int    `json:"quantity"`
	CustomerId string `json:"customerId"`
}
type CheckOut struct {
	Image       []byte  `json:"image"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Total       float64 `json:"total"`
	ItemId      string  `json:"itemId"`
}
type CheckOutHelper struct {
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Total       float64 `json:"total"`
	ItemId      string  `json:"itemId"`
}
