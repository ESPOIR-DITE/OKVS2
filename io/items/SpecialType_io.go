package items

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const specialtypeURL = api.BASE_URL + "specialType/"

func CreateSpecialType(obj items.SpecialType) (items.SpecialType, error) {
	entity := items.SpecialType{}
	resp, _ := api.Rest().SetBody(obj).Post(specialtypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetSpecialType(id string) (items.SpecialType, error) {
	entity := items.SpecialType{}
	resp, _ := api.Rest().Get(specialtypeURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteSpecialType(id string) (items.SpecialType, error) {
	entity := items.SpecialType{}
	resp, _ := api.Rest().Get(specialtypeURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetSpecialTypes() ([]items.SpecialType, error) {
	entity := []items.SpecialType{}
	resp, _ := api.Rest().Get(specialtypeURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateSpecialType(obj items.SpecialType) (items.SpecialType, error) {
	entity := items.SpecialType{}
	resp, _ := api.Rest().SetBody(obj).Post(specialtypeURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
