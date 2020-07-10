package brand

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const itemBraindURL = api.BASE_URL + "itemBraind"

type ItemBraind items.ItemBrand

func GetItemBrainds() ([]items.ItemBrand, error) {
	entities := []items.ItemBrand{}
	resp, _ := api.Rest().Get(itemBraindURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetItemBraind(id string) (items.ItemBrand, error) {
	entity := items.ItemBrand{}
	resp, _ := api.Rest().Get(itemBraindURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func CreateItemBraind(braind string) (items.ItemBrand, error) {
	fmt.Println(" we are about to creating Color", braind)
	entity := items.ItemBrand{}
	myType := Braind{"000", braind}
	resp, _ := api.Rest().SetBody(myType).Post(itemBraindURL + "/create")
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
func DeleteItemBraind(braind string) (items.ItemBrand, error) {
	//entities:=[]Color{}
	entity := items.ItemBrand{}
	resp, _ := api.Rest().Get(itemBraindURL + "/delete?id=" + braind)
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
func UpdateItemBraind(entit items.ItemBrand) (items.ItemBrand, error) {
	entity := items.ItemBrand{}
	resp, _ := api.Rest().SetBody(entit).Post(itemBraindURL + "/update")
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
	resp, _ := api.Rest().Get(itemBraindURL + "/readWithItemId?itemId=" + itemId)
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
