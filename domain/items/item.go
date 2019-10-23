package items

type BeautyMakeup struct {

	//ItemNumber string `json:"item_number"`
	ItemName   string `json:"item_name"`
	Size       string `json:"size"`
	Decription string `json:"decription"`
	Color      string `json:"color"`
}

type Cloths struct {
	ItemNumber string `json:"item_number"`
	ItemName   string `json:"item_name"`
	ItemType   string `json:"item_type"`
	Size       string `json:"size"`
	Marque     string `json:"marque"`
	Gender     string `json:"gender"`
	Decription string `json:"decription"`
	Color      string `json:"color"`
}
type Hair struct {
	ItemNumber string `json:"item_number"`
	Itemname   string `json:"itemname"`
	ItemType   string `json:"item_type"`
	Size       string `json:"size"`
	Decription string `json:"decription"`
	Color      string `json:"color"`
}
type Shoes struct {
	ItemNumber string `json:"itemNumber"`
	Marque     string `json:"marque"`
	ItemType   string `json:"itemtype"`
	Size       string `json:"size"`
	Gender     string `json:"gender"`
	Decription string `json:"decription"`
	Color      string `json:"color"`
}
type Items struct {
	ItemNumber  string `json:"item_number"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
	Description string `json:"description"`
}
type ItemSold struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id"`
	OrderId string `json:"order_id"`
}
type BeautyHelper struct {
	ItemName   string `json:"itemName"`
	Size       string `json:"size"`
	Decription string `json:"decription"`
	Color      string `json:"color"`
	Image      []byte `json:"image"`
}
