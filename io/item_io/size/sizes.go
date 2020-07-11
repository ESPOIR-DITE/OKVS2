package size

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const sizeURL = api.BASE_URL + "size"

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
func CreateSize(itemSize items.Size) (items.Size, error) {
	entity := items.Size{}
	resp, _ := api.Rest().SetBody(itemSize).Post(sizeURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
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
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
