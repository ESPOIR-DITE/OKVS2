package _type

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const itemTypeURL = api.BASE_URL + "itemtype"

func GetProductTypes() ([]items.ItemType, error) {
	//entity :=Color{}
	entities := []items.ItemType{}
	resp, _ := api.Rest().Get(itemTypeURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetProductType(id string) (items.ItemType, error) {
	entity := items.ItemType{}
	resp, _ := api.Rest().Get(itemTypeURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	fmt.Println("In Product TypeOfItem", entity)
	return entity, nil
}
func GetAllOfProductType(id string) ([]items.ItemType, error) {
	entity := []items.ItemType{}
	resp, _ := api.Rest().Get(itemTypeURL + "/readAll?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	fmt.Println("In Product TypeOfItem", entity)
	return entity, nil
}

func CreateProductType(itemType items.ItemType) (items.ItemType, error) {
	entity := items.ItemType{}
	resp, _ := api.Rest().SetBody(itemType).Post(itemTypeURL + "/create")
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
func DeleteProductType(productId string) (items.ItemType, error) {
	entity := items.ItemType{}
	resp, _ := api.Rest().Get(itemTypeURL + "/delete?id=" + productId)
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
