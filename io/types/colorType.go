package types

import (
	"OKVS2/api"

	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const colorURL = api.BASE_URL + "color"

type Color items.Color

func GetColors() ([]Color, error) {
	//entity :=Color{}
	entities := []Color{}
	resp, _ := api.Rest().Get(colorURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func CreateColors(color string) (Color, error) {
	fmt.Println(" we are about to creating Color", color)
	entity := Color{}
	myType := Color{"000", color}
	resp, _ := api.Rest().SetBody(myType).Post(colorURL + "/create")
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
func DeleteColor(color string) (items.Color, error) {
	//entities:=[]Color{}
	entity := items.Color{}
	resp, _ := api.Rest().Get(colorURL + "/delete?id=" + color)
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
