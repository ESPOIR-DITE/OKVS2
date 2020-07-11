package color

import (
	"OKVS2/api"

	"OKVS2/domain/items"
	"errors"
)

const colorURL = api.BASE_URL + "color"

type Color items.Color

func GetColors() ([]items.Color, error) {
	entities := []items.Color{}
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
func GetColor(id string) (items.Color, error) {
	entity := items.Color{}
	resp, _ := api.Rest().Get(colorURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateColors(color string) (Color, error) {
	entity := Color{}
	myType := Color{"000", color}
	resp, _ := api.Rest().SetBody(myType).Post(colorURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		//fmt.Println(" erro when marshaling", err)
		return entity, errors.New(resp.Status())

	}
	return entity, nil
}
func UpdateColors(color string) (Color, error) {
	entity := Color{}
	myType := Color{"000", color}
	resp, _ := api.Rest().SetBody(myType).Post(colorURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		//fmt.Println(" erro when marshaling", err)
		return entity, errors.New(resp.Status())

	}
	return entity, nil
}
func DeleteColor(color string) (items.Color, error) {
	entity := items.Color{}
	resp, _ := api.Rest().Get(colorURL + "/delete?id=" + color)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
