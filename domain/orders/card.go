package orders

type Card struct {
	Id         string `json:"id"`
	ItemId     string `json:"itemId"`
	CustomerId string `json:"customerId"`
	Quantity   int    `json:"quantity"`
}
