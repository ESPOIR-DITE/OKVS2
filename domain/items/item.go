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
	Itemtype   string   `json:"itemType"`
	Quantity   int      `json:"quantity"`
	Price      float64  `json:"price"`
	Image      [][]byte `json:"image"`
	Color      []string `json:"colors"`
	Braind     string   `json:"braind"`
}
type Images struct {
	Id    string `json:"id"`
	Image []byte `json:"image"`
}
type Item_Pictures struct {
	Id      string `json:"id"`
	ItemId  string `json:"itemId"`
	ImageId string `json:"imageId"`
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
	ItemId  string `json:"itemId"`
	ColorId string `json:"colorId"`
}
type Braind struct {
	BraindId   string `json:"braindId"`
	BraindName string `json:"braindName"`
}
type Accounting struct {
	ItemId   string  `json:"itemId"`
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
	SizeNumber string `json:"sizeNumber"`
}
type ProductSize struct {
	ItemId string `json:"itemId"`
	SizeId string `json:"sizeId"`
}
type AddressType struct {
	AddressTypeId string `json:"addressTypeId"`
	AddressType   string `json:"addressType"`
}
type ProductType struct {
	ItemId string `json:"itemId"`
	TypeId string `json:"typeId"`
}
type ItemType struct {
	ItemId string `json:"item_id"`
	TypeId string `json:"type_id"`
}
type Type struct {
	Id       string `json:"id"`
	TypeName string `json:"typeName"`
}
type ItemBraind struct {
	BraindId string `json:"braindId"`
	ItemId   string `json:"itemId"`
}
type ItemGender struct {
	ItemId   string `json:"itemId"`
	GenderId string `json:"genderId"`
}
type ItemView struct {
	ItemNumber  string  `json:"itemNumber"`
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       []byte  `json:"image"`
}
type ItemViewHtml struct {
	ItemNumber  string  `json:"itemNumber"`
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

//this entity help to read individua product
type ViewProduct struct {
	ItemId      string   `json:"itemId"`
	ItemName    string   `json:"itemName"`
	ItemBrand   string   `json:"itemBrand"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
	Pictures    [][]byte `json:"pictures"`
	Colors      []Color  `json:"colors"`
}

//this entity help ViewProduct
type ViewProduct2 struct {
	ItemId      string  `json:"itemId"`
	ItemName    string  `json:"itemName"`
	ItemBrand   string  `json:"itemBrand"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Colors      []Color `json:"colors"`
}
type Specials struct {
}
