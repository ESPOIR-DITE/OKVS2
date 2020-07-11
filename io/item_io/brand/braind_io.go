package brand

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const brandURL = api.BASE_URL + "brand"

func CreateBraind(braind string) (items.Brand, error) {
	entity := items.Brand{}
	resp, _ := api.Rest().SetBody(entity).Post(brandURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateBraind(braind string) (items.Brand, error) {
	entity := items.Brand{}
	resp, _ := api.Rest().SetBody(entity).Post(brandURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetBrainds() ([]items.Brand, error) {
	entities := []items.Brand{}
	resp, _ := api.Rest().Get(brandURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetBrand(id string) (items.Brand, error) {
	entity := items.Brand{}
	resp, _ := api.Rest().Get(brandURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteBraind(braind string) (items.Brand, error) {
	entity := items.Brand{}
	resp, _ := api.Rest().Get(brandURL + "/delete?id=" + braind)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
