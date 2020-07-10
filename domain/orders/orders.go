package orders

import "time"

type Card struct {
	Id         string `json:"id"`
	ItemId     string `json:"itemId"`
	CustomerId string `json:"customerId"`
	Quantity   int    `json:"quantity"`
}

type OrderHelper struct {
	OrderId   string  `json:"orderId"`
	Date      string  `json:"date"`
	ItemName  string  `json:"itemName"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	OrderStat string  `json:"orderStat"`
}

type OrderLine struct {
	LineCode    string  `json:"lineCode"`
	ItemNumber  string  `json:"itemNumber"`
	OrderNumber string  `json:"orderNumber"`
	Quantity    float64 `json:"quantity"`
}

type Orders struct {
	Id         string `json:"id"`
	CustomerId string `json:"customerId"`
	Date       string `json:"date"`
}
type OrderStatus struct {
	Id         string    `json:"id"`
	OrderId    string    `json:"orderId"`
	Date       time.Time `json:"date"`
	ModifiedBy string    `json:"modifiedBy"`
	Stat       string    `json:"stat"`
}

type Status struct {
	Id   string `json:"id"`
	Stat string `json:"stat"`
}

//type OrderHelper struct {
//	ItemNumber string `json:"itemNumber"`
//	Quantity   int    `json:"quantity"`
//	CustomerId string `json:"customerId"`
//}
