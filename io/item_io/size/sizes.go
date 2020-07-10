package size

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const sizeURL = api.BASE_URL + "size"

type Size items.Size

func GetSizes() ([]items.Size, error) {
	//entity :=Color{}
	entities := []items.Size{}
	resp, _ := api.Rest().Get(sizeURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetSize(id string) (items.Size, error) {
	entity := items.Size{}
	resp, _ := api.Rest().Get(sizeURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateSize(braind string) (Size, error) {
	fmt.Println(" we are about to creating Color", braind)
	entity := Size{}
	myType := Size{"000", braind}
	resp, _ := api.Rest().SetBody(myType).Post(sizeURL + "/create")
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
func DeleteSize(braind string) (items.Size, error) {
	//entities:=[]Color{}
	entity := items.Size{}
	resp, _ := api.Rest().Get(sizeURL + "/delete?id=" + braind)
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
