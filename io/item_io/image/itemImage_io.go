package image

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const itemImageURL = api.BASE_URL + "itemPicture"

func CreateItemImage(item items.ItemImage) (items.ItemImage, error) {
	entity := items.ItemImage{}
	resp, _ := api.Rest().SetBody(item).Post(itemImageURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateItemImage(image items.ItemImage) (items.ItemImage, error) {
	entity := items.ItemImage{}
	resp, _ := api.Rest().SetBody(image).Post(itemImageURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetItemImage(id string) (items.ItemImage, error) {
	entity := items.ItemImage{}
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

func GetItemImages() ([]items.ItemImage, error) {
	entities := []items.ItemImage{}
	resp, _ := api.Rest().Get(itemImageURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func DeleteItemImage(braind string) (items.ItemImage, error) {
	entity := items.ItemImage{}
	resp, _ := api.Rest().Get(itemImageURL + "/delete?id=" + braind)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
