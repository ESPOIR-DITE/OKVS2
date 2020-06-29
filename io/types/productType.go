package types

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const productTypeURL = api.BASE_URL + "productType"

type ProductType items.ProductType

func GetProductTypes() ([]items.ProductType, error) {
	//entity :=Color{}
	entities := []items.ProductType{}
	resp, _ := api.Rest().Get(productTypeURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetProductType(id string) (items.ProductType, error) {
	entity := items.ProductType{}
	resp, _ := api.Rest().Get(productTypeURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	fmt.Println("In Product Types", entity)
	return entity, nil
}

func GetAllOfProductType(id string) ([]items.ProductType, error) {
	entity := []items.ProductType{}
	resp, _ := api.Rest().Get(productTypeURL + "/readAll?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	fmt.Println("In Product Types", entity)
	return entity, nil
}

/***
!!!! can not create !!!!
*/
func CreateProductType(productName, desc string) (items.ProductType, error) {
	fmt.Println(" we are about to creating type", productName)
	entity := items.ProductType{}
	myType := items.Products{"000", productName, desc}
	resp, _ := api.Rest().SetBody(myType).Post(productTypeURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	fmt.Println(" we have create product", entity)
	if err != nil {
		fmt.Println(" erro when marshaling", err)
		return entity, errors.New(resp.Status())

	}
	return entity, nil
}
func DeleteProductType(productId string) (items.ProductType, error) {
	//entities:=[]Color{}
	entity := items.ProductType{}
	resp, _ := api.Rest().Get(productTypeURL + "/delete?id=" + productId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	fmt.Println(" we are Deleting product", entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
