package accountting_io

import (
	"OKVS2/api"
	"OKVS2/domain/accounting"
	"errors"
)

const accountingURL = api.BASE_URL + "accounting"

func CreateAccounting(account accounting.Accounting) (accounting.Accounting, error) {
	entity := accounting.Accounting{}
	resp, _ := api.Rest().SetBody(account).Post(accountingURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateAccounting(account accounting.Accounting) (accounting.Accounting, error) {
	entity := accounting.Accounting{}
	resp, _ := api.Rest().SetBody(account).Post(accountingURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetAccountings() ([]accounting.Accounting, error) {
	//entity :=Color{}
	entities := []accounting.Accounting{}
	resp, _ := api.Rest().Get(accountingURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetAccounting(id string) (accounting.Accounting, error) {
	entity := accounting.Accounting{}
	resp, _ := api.Rest().Get(accountingURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func DeleteAccounting(id string) (accounting.Accounting, error) {
	entity := accounting.Accounting{}
	resp, _ := api.Rest().Get(accountingURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
