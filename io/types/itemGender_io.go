package types

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const itemGenderURL = api.BASE_URL + "itemGender"

type ItemGender items.ItemGender

func GetItemGenders() ([]items.ItemGender, error) {
	//entity :=Color{}
	entities := []items.ItemGender{}
	resp, _ := api.Rest().Get(itemGenderURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetItemGender(id string) (items.ItemGender, error) {
	entity := items.ItemGender{}
	resp, _ := api.Rest().Get(itemGenderURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	fmt.Println(" we are about to reading ItemGender <<int item gender>>", entity)
	return entity, nil

}
func CreateItemGender(braind string) (items.ItemGender, error) {
	fmt.Println(" we are about to creating Color", braind)
	entity := items.ItemGender{}
	myType := Braind{"000", braind}
	resp, _ := api.Rest().SetBody(myType).Post(itemGenderURL + "/create")
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
func DeleteItemGender(braind string) (items.ItemGender, error) {
	//entities:=[]Color{}
	entity := items.ItemGender{}
	resp, _ := api.Rest().Get(itemGenderURL + "/delete?id=" + braind)
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
func ReadItemGenderWithItemId(itemId string) (items.ItemGender, error) {
	entity := items.ItemGender{}
	resp, _ := api.Rest().Get(itemGenderURL + "/readWithItemId?itemId=" + itemId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateItemGender(itemGender items.ItemGender) (items.ItemGender, error) {
	entity := items.ItemGender{}
	resp, _ := api.Rest().SetBody(itemGender).Post(itemGenderURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
