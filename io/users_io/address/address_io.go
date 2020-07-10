package users_io

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
)

const addressURL = api.BASE_URL + "address"

func CreateAddress(obj users.UserAddress) (users.Address, error) {
	entity := users.Address{}
	resp, _ := api.Rest().SetBody(obj).Post(addressURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetAddress(id string) (users.Address, error) {
	entity := users.Address{}
	resp, _ := api.Rest().Get(addressURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetAddresss() ([]users.Address, error) {
	entity := []users.Address{}
	resp, _ := api.Rest().Get(addressURL + "/reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteAddress(id string) (users.Address, error) {
	entity := users.Address{}
	resp, _ := api.Rest().Get(addressURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateAddress(obj users.Address) (users.Address, error) {
	entity := users.Address{}
	resp, _ := api.Rest().SetBody(obj).Post(addressURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
