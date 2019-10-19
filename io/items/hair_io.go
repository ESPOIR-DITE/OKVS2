package items

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const hairURL = api.BASE_URL + "/hair"

type HairItem items.Hair

func GetHairs() ([]HairItem, error) {
	entities := []HairItem{}
	resp, _ := api.Rest().Get(hairURL + "/reads")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)

	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetHair(id string) (HairItem, error) {
	entity := HairItem{}

	resp, _ := api.Rest().Get(hairURL + "/read" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateHair(cloth interface{}) (HairItem, error) {
	entity := HairItem{}

	resp, _ := api.Rest().SetBody(cloth).Post(hairURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteHair(id string) (HairItem, error) {
	entity := HairItem{}
	resp, _ := api.Rest().Get(hairURL + "delete" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateHair(beauty interface{}) (HairItem, error) {
	entity := HairItem{}
	resp, _ := api.Rest().SetBody(beauty).Post(hairURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
