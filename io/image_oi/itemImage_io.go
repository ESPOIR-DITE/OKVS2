package image_oi

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

type ItemImage items.Item_Pictures

const itemImageURL = api.BASE_URL + "image"

func GetItemImage(id string) ([]items.Images, error) {
	entity := []items.Images{}

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
