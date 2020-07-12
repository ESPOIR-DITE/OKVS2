package users_io

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
	"fmt"
)

const addressTypeURL = api.BASE_URL + "addressType"

func GetAddressTypes() ([]users.AddressType, error) {
	entities := []users.AddressType{}
	resp, _ := api.Rest().Get(addressTypeURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func CreateAddressType(addressT string) (users.AddressType, error) {
	entity := users.AddressType{}
	resp, _ := api.Rest().SetBody(addressT).Post(addressTypeURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteAddressType(addressT string) (users.AddressType, error) {
	entity := users.AddressType{}
	resp, _ := api.Rest().Get(addressTypeURL + "/delete?id=" + addressT)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWithAddressType(typeName string) (users.AddressType, error) {
	entity := users.AddressType{}
	resp, _ := api.Rest().Get(addressTypeURL + "/readwithType?id=" + typeName)
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
