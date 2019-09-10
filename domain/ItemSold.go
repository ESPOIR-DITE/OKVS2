package domain

type ItemSold struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id"`
	OrderId string `json:"order_id"`
}
