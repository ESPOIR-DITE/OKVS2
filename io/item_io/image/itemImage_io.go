package image

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

type ItemImage items.ItemImage

const itemImageURL = api.BASE_URL + "itemPicture"

func GetItemImage(id string) ([]items.ItemImage, error) {
	entity := []items.ItemImage{}

	resp, _ := api.Rest().Get(itemImageURL + "/read?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWithImageId(imageId string) (items.ItemImage, error) {
	entity := items.ItemImage{}
	resp, _ := api.Rest().Get(itemImageURL + "/readWithImageId?id=" + imageId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
