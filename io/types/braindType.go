package types

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const braindURL = api.BASE_URL + "braind"

type Braind items.Braind

func GetBrainds() ([]Braind, error) {
	//entity :=Color{}
	entities := []Braind{}
	resp, _ := api.Rest().Get(braindURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func CreateBraind(braind string) (Braind, error) {
	fmt.Println(" we are about to creating Color", braind)
	entity := Braind{}
	myType := Braind{"000", braind}
	resp, _ := api.Rest().SetBody(myType).Post(braindURL + "/create")
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
func DeleteBraind(braind string) (items.Braind, error) {
	//entities:=[]Color{}
	entity := items.Braind{}
	resp, _ := api.Rest().Get(braindURL + "/delete?id=" + braind)
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
