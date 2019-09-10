package controller

type OrderLine struct {
	LineId  string `json:"line_id"`
	ItemId  string `json:"item_id"`
	OrderId string `json:"order_id"`
	Date    string `json:"date"`
}
