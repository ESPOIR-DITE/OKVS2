package accountting_io

import (
	"OKVS2/api"
	"OKVS2/domain/accounting"
	"errors"
)

const accountchangeURL = api.BASE_URL + "accountchange"

func CreateAccountChange(account accounting.AccountChange) (accounting.AccountChange, error) {
	entity := accounting.AccountChange{}
	resp, _ := api.Rest().SetBody(account).Post(accountchangeURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateAccountChange(account accounting.AccountChange) (accounting.AccountChange, error) {
	entity := accounting.AccountChange{}
	resp, _ := api.Rest().SetBody(account).Post(accountchangeURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetAccountChanges() ([]accounting.AccountChange, error) {
	//entity :=Color{}
	entities := []accounting.AccountChange{}
	resp, _ := api.Rest().Get(accountchangeURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetAccountChange(id string) (accounting.AccountChange, error) {
	entity := accounting.AccountChange{}
	resp, _ := api.Rest().Get(accountchangeURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func DeleteAccountChange(id string) (accounting.AccountChange, error) {
	entity := accounting.AccountChange{}
	resp, _ := api.Rest().Get(accountchangeURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
