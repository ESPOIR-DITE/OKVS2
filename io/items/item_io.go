package items

import (
	"OKVS2/api"
	ItemValeu "OKVS2/domain/items"
	"errors"
)

const itemURL = api.BASE_URL + "/item"

type Item ItemValeu.Items

func GetItems() ([]Item, error) {
	entities := []Item{}
	resp, _ := api.Rest().Get(itemURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetItem(id string) (Item, error) {
	entity := Item{}
	resp, _ := api.Rest().Get(itemURL + "/read" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateItem(item Item) (Item, error) {
	entity := Item{}
	resp, _ := api.Rest().SetBody(item).Post(itemURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteItem(id string) (Item, error) {
	entity := Item{}
	resp, _ := api.Rest().Get(itemURL + "/delete" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateItem(item interface{}) (Item, error) {
	entity := Item{}
	resp, _ := api.Rest().SetBody(item).Post(itemURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
