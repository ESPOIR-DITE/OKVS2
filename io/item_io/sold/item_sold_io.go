package sold

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const itemsoldURL = api.BASE_URL + "itemsold"

func CreateSold(braind string) (items.ItemSold, error) {
	entity := items.ItemSold{}
	resp, _ := api.Rest().SetBody(entity).Post(itemsoldURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateSold(braind string) (items.ItemSold, error) {
	entity := items.ItemSold{}
	resp, _ := api.Rest().SetBody(entity).Post(itemsoldURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetSolds() ([]items.ItemSold, error) {
	entities := []items.ItemSold{}
	resp, _ := api.Rest().Get(itemsoldURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetSold(id string) (items.ItemSold, error) {
	entity := items.ItemSold{}
	resp, _ := api.Rest().Get(itemsoldURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteSold(braind string) (items.ItemSold, error) {
	entity := items.ItemSold{}
	resp, _ := api.Rest().Get(itemsoldURL + "/delete?id=" + braind)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
