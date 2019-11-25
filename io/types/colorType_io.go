package types

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const itemColorURL = api.BASE_URL + "itemColor"

type ItemColor items.ItemColor

func GetItemColors() ([]items.ItemColor, error) {
	//entity :=Color{}
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

/**
func CreateColors(color string) (Color, error) {
	fmt.Println(" we are about to creating Color", color)
	entity := Color{}
	myType := Color{"000", color}
	resp, _ := api.Rest().SetBody(myType).Post(colorURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	fmt.Println(" we have create Color", entity)
	if err != nil {
		fmt.Println(" erro when marshaling", err)
		return entity, errors.New(resp.Status())

	}
	return entity, nil
}
func DeleteColor(color string) (items.Color, error) {
	//entities:=[]Color{}
	entity := items.Color{}
	resp, _ := api.Rest().Get(colorURL + "/delete?id=" + color)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	fmt.Println(" we are Deleting Color", entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}*/
