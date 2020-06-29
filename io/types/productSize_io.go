package types

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const productSizeURL = api.BASE_URL + "productSize"

type ProductSize items.ProductSize

func GetProductSizes() ([]items.ProductSize, error) {
	//entity :=Color{}
	entities := []items.ProductSize{}
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
func GetProductSize(id string) (items.ProductSize, error) {
	entity := items.ProductSize{}
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
func GetPtoductSizeWithItemId(itemId string) ([]items.ProductSize, error) {
	entity := []items.ProductSize{}
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
func CreateProductSize(braind string) (items.ProductSize, error) {
	fmt.Println(" we are about to creating Color", braind)
	entity := items.ProductSize{}
	myType := Braind{"000", braind}
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
func DeleteProductSize(id string) (items.ProductSize, error) {
	//entities:=[]Color{}
	entity := items.ProductSize{}
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
func CreateAllProductSize(sizeIdList []items.ProductSize) (bool, error) {
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
