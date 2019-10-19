package items

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const clothsURL = api.BASE_URL + "/cloths"

type ClothsItem items.Cloths

func GetCloths() ([]ClothsItem, error) {
	entities := []ClothsItem{}
	resp, _ := api.Rest().Get(clothsURL + "/reads")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)

	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetCloth(id string) (ClothsItem, error) {
	entity := ClothsItem{}

	resp, _ := api.Rest().Get(clothsURL + "/read" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateCloths(cloth interface{}) (ClothsItem, error) {
	entity := ClothsItem{}

	resp, _ := api.Rest().SetBody(cloth).Post(clothsURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCloths(id string) (ClothsItem, error) {
	entity := ClothsItem{}
	resp, _ := api.Rest().Get(clothsURL + "delete" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCloths(beauty interface{}) (ClothsItem, error) {
	entity := ClothsItem{}
	resp, _ := api.Rest().SetBody(beauty).Post(clothsURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
