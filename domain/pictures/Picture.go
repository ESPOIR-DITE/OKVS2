package pictures

import "os"

type MypicHelper struct {
	ItemId      string  `json:"item_id"`
	Image       os.File `json:"image"`
	Description string  `json:"description"`
}
type Mypic struct {
	Id          string `json:"id"`
	ItemId      string `json:"item_id"`
	Description string `json:"description"`
}
