package items

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const specialURL = api.BASE_URL + "special/"

func CreateSpecial(obj items.Specials) (items.Specials, error) {
	entity := items.Specials{}
	resp, _ := api.Rest().SetBody(obj).Post(specialURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetSpecial(specialId string) (items.Specials, error) {
	entity := items.Specials{}
	resp, _ := api.Rest().Get(specialURL + "read?=id" + specialId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteSpecial(specialId string) (items.Specials, error) {
	entity := items.Specials{}
	resp, _ := api.Rest().Get(specialURL + "delete?=id" + specialId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetAllSpecial() ([]items.Specials, error) {
	entity := []items.Specials{}
	resp, _ := api.Rest().Get(specialURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateSpecial(obj items.Specials) (items.Specials, error) {
	entity := items.Specials{}
	resp, _ := api.Rest().SetBody(obj).Post(specialURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
