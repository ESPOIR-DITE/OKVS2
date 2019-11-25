package types

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const itemBraindURL = api.BASE_URL + "itemBraind"

type ItemBraind items.ItemBraind

func GetItemBrainds() ([]items.Braind, error) {
	//entity :=Color{}
	entities := []items.Braind{}
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
func GetItemBraind(id string) (items.Braind, error) {
	entity := items.Braind{}
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
func CreateItemBraind(braind string) (items.Braind, error) {
	fmt.Println(" we are about to creating Color", braind)
	entity := items.Braind{}
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
func DeleteItemBraind(braind string) (items.Braind, error) {
	//entities:=[]Color{}
	entity := items.Braind{}
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
