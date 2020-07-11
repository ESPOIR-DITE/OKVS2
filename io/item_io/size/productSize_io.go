package size

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const itemSizeURL = api.BASE_URL + "itemSize"

func GetProductSizes() ([]items.ItemSize, error) {
	entities := []items.ItemSize{}
	resp, _ := api.Rest().Get(itemSizeURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetProductSize(id string) (items.ItemSize, error) {
	entity := items.ItemSize{}
	resp, _ := api.Rest().Get(itemSizeURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func GetPtoductSizeWithItemId(itemId string) ([]items.ItemSize, error) {
	entity := []items.ItemSize{}
	resp, _ := api.Rest().Get(itemSizeURL + "/readWithItem?id=" + itemId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateProductSize(itemSize items.ItemSize) (items.ItemSize, error) {
	entity := items.ItemSize{}
	resp, _ := api.Rest().SetBody(itemSize).Post(itemSizeURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteProductSize(id string) (items.ItemSize, error) {
	entity := items.ItemSize{}
	resp, _ := api.Rest().Get(itemSizeURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteAllOfProductSize(projectId string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(itemSizeURL + "/deleteAllOf?productId=" + projectId)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func CreateAllProductSize(sizeIdList []items.ItemSize) (bool, error) {
	var entity bool
	resp, _ := api.Rest().SetBody(sizeIdList).Post(itemSizeURL + "/createAll")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
