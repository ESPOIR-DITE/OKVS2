package customer

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
)

const customerURL = api.BASE_URL + "/customer"

type Customer users.Customer

func GetCustomers() ([]users.Customer, error) {
	entities := []users.Customer{}
	resp, _ := api.Rest().Get(customerURL + "/reads")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetCustomer(id string) (users.Customer, error) {
	entity := users.Customer{}

	reso, _ := api.Rest().Get(customerURL + "/read?id=" + id)

	if reso.IsError() {
		return entity, errors.New(reso.Status())
	}
	err := api.JSON.Unmarshal(reso.Body(), &entity)
	if err != nil {
		return entity, errors.New(reso.Status())

	}
	return entity, nil
}
func CreateCustomer(entit interface{}) (Customer, error) {
	entity := Customer{}

	resp, _ := api.Rest().SetBody(entit).Post(customerURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())

	}
	return entity, nil
}
func DeleteAdmin(id string) (Customer, error) {
	entity := Customer{}
	resp, _ := api.Rest().Get(customerURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCustomer(entity interface{}) (bool, error) {
	resp, _ := api.Rest().SetBody(entity).Post(customerURL + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
