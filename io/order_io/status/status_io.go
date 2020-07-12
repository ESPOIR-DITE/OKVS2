package status

import (
	"OKVS2/api"
	"OKVS2/domain/orders"
	"errors"
)

const statusURL = api.BASE_URL + "status/"

func CreateStatus(obj orders.Status) (orders.Status, error) {
	entities := orders.Status{}
	resp, _ := api.Rest().SetBody(obj).Post(statusURL + "create")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func UpdateStatus(obj orders.Status) (orders.Status, error) {
	entities := orders.Status{}
	resp, _ := api.Rest().SetBody(obj).Post(statusURL + "update")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetStatus(id string) (orders.Status, error) {
	entities := orders.Status{}
	resp, _ := api.Rest().Get(statusURL + "read?id=" + id)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func DeleteStatus(id string) (orders.Status, error) {
	entities := orders.Status{}
	resp, _ := api.Rest().Get(statusURL + "delete?id=" + id)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetStatues() ([]orders.Status, error) {
	entities := []orders.Status{}
	resp, _ := api.Rest().Get(statusURL + "reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
