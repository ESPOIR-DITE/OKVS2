package brand

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const itemBrandURL = api.BASE_URL + "itemBrand"

func CreateItemBrand(itemBrand items.ItemBrand) (items.ItemBrand, error) {
	entity := items.ItemBrand{}
	resp, _ := api.Rest().SetBody(itemBrand).Post(itemBrandURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetItemBrand(id string) (items.ItemBrand, error) {
	entity := items.ItemBrand{}
	resp, _ := api.Rest().Get(itemBrandURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetItemBrands() ([]items.ItemBrand, error) {
	entities := []items.ItemBrand{}
	resp, _ := api.Rest().Get(itemBrandURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func DeleteItemBrand(braind string) (items.ItemBrand, error) {
	entity := items.ItemBrand{}
	resp, _ := api.Rest().Get(itemBrandURL + "/delete?id=" + braind)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateItemBrand(entit items.ItemBrand) (items.ItemBrand, error) {
	entity := items.ItemBrand{}
	resp, _ := api.Rest().SetBody(entit).Post(itemBrandURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWithItemId(itemId string) (items.ItemBrand, error) {
	entity := items.ItemBrand{}
	resp, _ := api.Rest().Get(itemBrandURL + "/readWithItemId?itemId=" + itemId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
