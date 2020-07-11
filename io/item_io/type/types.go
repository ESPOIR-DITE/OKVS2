package _type

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const typesURL = api.BASE_URL + "type"

func CreateType(typeOfItem items.TypeOfItem) (items.TypeOfItem, error) {
	entity := items.TypeOfItem{}
	resp, _ := api.Rest().SetBody(typeOfItem).Post(typesURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateType(typeOfItem items.TypeOfItem) (items.TypeOfItem, error) {
	entity := items.TypeOfItem{}
	resp, _ := api.Rest().SetBody(typeOfItem).Post(typesURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetType() (items.TypeOfItem, error) {
	entities := items.TypeOfItem{}
	resp, _ := api.Rest().Get(typesURL + "/read")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetTypes() ([]items.TypeOfItem, error) {
	entities := []items.TypeOfItem{}
	resp, _ := api.Rest().Get(typesURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func DeleteType(braind string) (items.TypeOfItem, error) {
	entity := items.TypeOfItem{}
	resp, _ := api.Rest().Get(typesURL + "/delete?id=" + braind)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
