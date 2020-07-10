package users_io

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const addressTypeURL = api.BASE_URL + "addressType"

type AddressType items.AddressType

func GetAddressTypes() ([]AddressType, error) {
	entities := []AddressType{}
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
func CreateAddressType(addressT string) (AddressType, error) {
	fmt.Println(" we are about to creating addressT", addressT)
	entity := AddressType{}
	myType := AddressType{"000", addressT}
	resp, _ := api.Rest().SetBody(myType).Post(addressTypeURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	//fmt.Println(" we have create addressT", entity)
	if err != nil {
		fmt.Println(" erro when marshaling", err)
		return entity, errors.New(resp.Status())

	}
	return entity, nil
}
func DeleteAddressType(addressT string) (items.AddressType, error) {
	//entities:=[]Color{}
	entity := items.AddressType{}
	resp, _ := api.Rest().Get(addressTypeURL + "/delete?id=" + addressT)
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
func ReadWithAddressType(typeName string) (items.AddressType, error) {
	entity := items.AddressType{}
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
