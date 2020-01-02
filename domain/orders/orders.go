package orders

import "time"

type Orders struct {
	Id         string `json:"id"`
	CustomerId string `json:"customerId"`
	Date       string `json:"date"`
}

type OrderLine struct {
	Linecode    string  `json:"linecode"`
	ItemNumber  string  `json:"itemNumber"`
	OrderNumber string  `json:"orderNumber"`
	Quantity    float64 `json:"quantity"`
}
type OrderStatus struct {
	Id         string    `json:"id"`
	OrderId    string    `json:"orderId"`
	Date       time.Time `json:"date"`
	ModifiedBy string    `json:"modifiedBy"`
	Stat       string    `json:"stat"`
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
type OrderHelper struct {
	OrderId   string  `json:"orderId"`
	Date      string  `json:"date"`
	ItemName  string  `json:"itemName"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	OrderStat string  `json:"orderStat"`
}

type Status struct {
	Id   string `json:"id"`
	Stat string `json:"stat"`
}
