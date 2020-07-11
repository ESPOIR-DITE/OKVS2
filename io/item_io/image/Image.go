package image

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const imageURL = api.BASE_URL + "image"

func CreateImage(braind string) (items.Images, error) {
	entity := items.Images{}
	resp, _ := api.Rest().SetBody(entity).Post(imageURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateImage(image items.Images) (items.Images, error) {
	entity := items.Images{}
	resp, _ := api.Rest().SetBody(image).Post(imageURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetImage(id string) (items.Images, error) {
	entity := items.Images{}
	resp, _ := api.Rest().Get(imageURL + "/read?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetImages() ([]items.Images, error) {
	entities := []items.Images{}
	resp, _ := api.Rest().Get(imageURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func DeleteImage(braind string) (items.Images, error) {
	entity := items.Images{}
	resp, _ := api.Rest().Get(imageURL + "/delete?id=" + braind)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
