package order

import (
	"OKVS2/api"
	"OKVS2/domain/orders"
	"errors"
)

const orderURL = api.BASE_URL + "/order"

type Order orders.Orders

func GetCustomers() ([]Order, error) {
	entities := []Order{}
	resp, _ := api.Rest().Get(orderURL + "/reads")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetCustomer(id string) (Order, error) {
	entity := Order{}
	reso, _ := api.Rest().Get(orderURL + "/read" + id)
	if reso.IsError() {
		return entity, errors.New(reso.Status())
	}
	return entity, nil
}
func CreateCustomer(entit interface{}) (Order, error) {
	entity := Order{}
	resp, _ := api.Rest().SetBody(entit).Post(orderURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteAdmin(id string) (Order, error) {
	entity := Order{}
	resp, _ := api.Rest().Get(orderURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCustomer(entit interface{}) (Order, error) {
	entity := Order{}
	resp, _ := api.Rest().SetBody(entit).Post(orderURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
