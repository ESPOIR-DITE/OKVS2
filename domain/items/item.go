package items

type BeautyMakeup struct {
	//ItemNumber string `json:"item_number"`
	ItemName   string `json:"item_name"`
	Size       string `json:"size"`
	Decription string `json:"decription"`
	Color      string `json:"color"`
}
type MyItemHelper struct {
	ItemName   string   `json:"itemName"`
	Size       []string `json:"size"`
	Decription string   `json:"decription"`
	Gender     []string `json:"gender"`
	Itemtype   string   `json:"itemtype"`
	Quantity   int      `json:"quantity"`
	Price      float64  `json:"price"`
	Image      [][]byte `json:"image"`
	Color      []string `json:"color"`
	Braind     string   `json:"braind"`
}
type MyImages struct {
	Image  []byte
	Image1 []byte
	Image2 []byte
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
	Name        string `json:"name"`
	TypeName    string `json:"type_name"`
	Description string `json:"description"`
}
type ItemSold struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id"`
	OrderId string `json:"order_id"`
}

type Color struct {
	ColorId   string `json:"colorId"`
	ColorName string `json:"colorName"`
}
type ItemColor struct {
	ItemId  string `json:"item_id"`
	ColorId string `json:"color_id"`
}
type Braind struct {
	BraindId   string `json:"braindId"`
	BraindName string `json:"braindName"`
}
type Accounting struct {
	ItemId   string  `json:"item_id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
type Products struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Size struct {
	Id         string `json:"id"`
	SizeNumber string `json:"size_number"`
}
type ProductSize struct {
	ItemId string `json:"item_id"`
	SizeId string `json:"size_id"`
}
type AddressType struct {
	AddressTypeId string `json:"addressTypeId"`
	AddressType   string `json:"addressType"`
}
type ProductType struct {
	ItemId   string `json:"item_id"`
	TypeName string `json:"type_name"`
}
type ItemType struct {
	ItemId string `json:"item_id"`
	TypeId string `json:"type_id"`
}
type Type struct {
	Id       string `json:"id"`
	TypeName string `json:"type_name"`
}
