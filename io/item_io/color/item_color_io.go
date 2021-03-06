package color

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const itemColorURL = api.BASE_URL + "itemColor"

func GetItemColors() ([]items.ItemColor, error) {
	entities := []items.ItemColor{}
	resp, _ := api.Rest().Get(itemColorURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetItemColor(id string) (items.ItemColor, error) {
	entity := items.ItemColor{}
	resp, _ := api.Rest().Get(itemColorURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetItemColorList(id string) ([]items.ItemColor, error) {
	entity := []items.ItemColor{}
	resp, _ := api.Rest().Get(itemColorURL + "/readsfor?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteAllOfItemColor(itemId string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(itemColorURL + "/deleteAllFor?itemId=" + itemId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateAllOfItemColors(color []items.ItemColor) (bool, error) {
	var entity bool
	resp, _ := api.Rest().SetBody(color).Post(itemColorURL + "/createAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteItemColor(color string) (items.ItemColor, error) {
	//entities:=[]Color{}
	entity := items.ItemColor{}
	resp, _ := api.Rest().Get(colorURL + "/delete?id=" + color)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
