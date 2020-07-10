package _type

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const typesURL = api.BASE_URL + "type"

type Type items.TypeOfItem

func GetTypes() ([]items.TypeOfItem, error) {
	//entity :=Color{}
	entities := []items.TypeOfItem{}
	resp, _ := api.Rest().Get(typesURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func CreateType(braind string) (Type, error) {
	fmt.Println(" we are about to creating Color", braind)
	entity := Type{}
	myType := Type{"000", braind}
	resp, _ := api.Rest().SetBody(myType).Post(typesURL + "/create")
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
func DeleteType(braind string) (items.TypeOfItem, error) {
	//entities:=[]Color{}
	entity := items.TypeOfItem{}
	resp, _ := api.Rest().Get(typesURL + "/delete?id=" + braind)
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
