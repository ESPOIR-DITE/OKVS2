package items

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const shoesURL = api.BASE_URL + "/shoes"

type ShoesItem items.Shoes

func GetShoes() ([]ShoesItem, error) {
	entities := []ShoesItem{}
	resp, _ := api.Rest().Get(shoesURL + "/reads")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)

	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetShoe(id string) (ShoesItem, error) {
	entity := ShoesItem{}

	resp, _ := api.Rest().Get(shoesURL + "/read" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateShoes(cloth interface{}) (ShoesItem, error) {
	entity := ShoesItem{}

	resp, _ := api.Rest().SetBody(cloth).Post(shoesURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteShoes(id string) (ShoesItem, error) {
	entity := ShoesItem{}
	resp, _ := api.Rest().Get(shoesURL + "delete" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateShoes(beauty interface{}) (ShoesItem, error) {
	entity := ShoesItem{}
	resp, _ := api.Rest().SetBody(beauty).Post(shoesURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
