package customer

import (
	"OKVS2/api"
	"OKVS2/domain/gender"
	"errors"
)

const customergenderURL = api.BASE_URL + "customerGender"

func CreateCustomerGender(obj gender.CustomerGender) (gender.CustomerGender, error) {
	entity := gender.CustomerGender{}
	resp, _ := api.Rest().SetBody(obj).Post(customergenderURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetCustomerGender(id string) (gender.CustomerGender, error) {
	entity := gender.CustomerGender{}
	resp, _ := api.Rest().Get(customergenderURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetCustomerGenders() ([]gender.CustomerGender, error) {
	entity := []gender.CustomerGender{}
	resp, _ := api.Rest().Get(customergenderURL + "/reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCustomerGender(id string) (gender.CustomerGender, error) {
	entity := gender.CustomerGender{}
	resp, _ := api.Rest().Get(customergenderURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCustomerGender(obj gender.CustomerGender) (gender.CustomerGender, error) {
	entity := gender.CustomerGender{}
	resp, _ := api.Rest().SetBody(obj).Post(customergenderURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
