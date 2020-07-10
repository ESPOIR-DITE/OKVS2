package items

import "os"

type Brand struct {
	BrandId   string `json:"brandId"`
	BrandName string `json:"brandName"`
}
type ItemBrand struct {
	Id      string `json:"id"`
	BrandId string `json:"braindId"`
	ItemId  string `json:"itemId"`
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

/***
Color
*/
type Color struct {
	ColorId   string `json:"colorId"`
	ColorName string `json:"colorName"`
}
type ItemColor struct {
	Id      string `json:"id"`
	ItemId  string `json:"itemId"`
	ColorId string `json:"colorId"`
}
type ItemGender struct {
	Id       string `json:"id"`
	ItemId   string `json:"item_id"`
	GenderId string `json:"gender_id"`
}

type ItemImage struct {
	Id      string `json:"id"`
	ItemId  string `json:"itemId"`
	ImageId string `json:"imageId"`
}
type Images struct {
	Id    string `json:"id"`
	Image []byte `json:"image"`
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

type ItemSold struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id"`
	OrderId string `json:"order_id"`
}

type ItemType struct {
	ItemId string `json:"itemId"`
	TypeId string `json:"typeId"`
}
type TypeOfItem struct {
	Id       string `json:"id"`
	TypeName string `json:"typeName"`
}

type Size struct {
	Id         string `json:"id"`
	SizeNumber string `json:"sizeNumber"`
}
type ItemSize struct {
	Id     string `json:"id"`
	ItemId string `json:"itemId"`
	SizeId string `json:"sizeId"`
}

type Specials struct {
	SpecialId     string  `json:"specialId"`
	Title         string  `json:"title"`
	ItemId        string  `json:"itemId"`
	SpecialTypeId string  `json:"specialTypeId"`
	Period        string  `json:"period"`
	EndPeriod     string  `json:"endPeriod"`
	Description   string  `json:"description"`
	ActualPrice   float64 `json:"actualPrice"`
}

type MypicHelper struct {
	ItemId      string  `json:"item_id"`
	Image       os.File `json:"image"`
	Description string  `json:"description"`
}

type SpecialType struct {
	Id          string `json:"id"`
	SpecialType string `json:"specialType"`
	Description string `json:"description"`
}

type Items struct {
	ItemNumber  string `json:"item_number"`
	Name        string `json:"name"`
	TypeName    string `json:"type_name"`
	Description string `json:"description"`
}

//TODO sorting these out
type MyImages struct {
	Image  []byte
	Image1 []byte
	Image2 []byte
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
type ViewItem struct {
	ItemId      string   `json:"itemId"`
	ItemName    string   `json:"itemName"`
	ItemBrand   string   `json:"itemBrand"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
	Pictures    [][]byte `json:"pictures"`
	Colors      []Color  `json:"colors"`
}

//this entity help ViewItem
type ViewItem2 struct {
	ItemId      string  `json:"itemId"`
	ItemName    string  `json:"itemName"`
	ItemBrand   string  `json:"itemBrand"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Colors      []Color `json:"colors"`
}
