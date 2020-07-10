package size

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"OKVS2/io/item_io/brand"
	"errors"
	"fmt"
)

const productSizeURL = api.BASE_URL + "productSize"

type ProductSize items.ItemSize

func GetProductSizes() ([]items.ItemSize, error) {
	//entity :=Color{}
	entities := []items.ItemSize{}
	resp, _ := api.Rest().Get(productSizeURL + "/reads")
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
	resp, _ := api.Rest().Get(productSizeURL + "/read?id=" + id)
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
	resp, _ := api.Rest().Get(productSizeURL + "/readWithItem?id=" + itemId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateProductSize(braind string) (items.ItemSize, error) {
	fmt.Println(" we are about to creating Color", braind)
	entity := items.ItemSize{}
	myType := brand.Braind{"000", braind}
	resp, _ := api.Rest().SetBody(myType).Post(productSizeURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	fmt.Println(" we have create Color", entity)
	if err != nil {
		fmt.Println(" erro when marshaling", err)
		return entity, errors.New(resp.Status())

	}
	return entity, nil
}
func DeleteProductSize(id string) (items.ItemSize, error) {
	//entities:=[]Color{}
	entity := items.ItemSize{}
	resp, _ := api.Rest().Get(productSizeURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	fmt.Println(" we are Deleting Color", entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteAllOfProductSize(projectId string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(productSizeURL + "/deleteAllOf?productId=" + projectId)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		fmt.Println(" erro when marshaling", err)
		return false, errors.New(resp.Status())

	}
	return true, nil
}
func CreateAllProductSize(sizeIdList []items.ItemSize) (bool, error) {
	var entity bool
	resp, _ := api.Rest().SetBody(sizeIdList).Post(productSizeURL + "/createAll")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}