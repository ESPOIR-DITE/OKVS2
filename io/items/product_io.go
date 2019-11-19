package items

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const productURL = api.BASE_URL + "product"

type Product items.Products

func GetProducts() ([]items.Products, error) {
	//entity :=Color{}
	entities := []items.Products{}
	resp, _ := api.Rest().Get(productURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetProduct(id string) (items.Products, error) {
	entity := items.Products{}
	resp, _ := api.Rest().Get(productURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	fmt.Println("In Product", entity)
	return entity, nil
}
func CreateProduct(productName, desc string) (items.Products, error) {
	fmt.Println(" we are about to creating product", productName)
	entity := items.Products{}
	myType := items.Products{"000", productName, desc}
	resp, _ := api.Rest().SetBody(myType).Post(productURL + "/create")
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
func DeleteProduct(productId string) (items.Products, error) {
	//entities:=[]Color{}
	entity := items.Products{}
	resp, _ := api.Rest().Get(productURL + "/delete?id=" + productId)
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
