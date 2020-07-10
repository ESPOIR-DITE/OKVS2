package accountting_io

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const accountingURL = api.BASE_URL + "acounting"

type Accounting items.Accounting

func GetAccounting(id string) (items.Accounting, error) {
	entity := items.Accounting{}
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
func UpdateAccounting(account items.Accounting) (items.Accounting, error) {
	entity := items.Accounting{}
	resp, _ := api.Rest().SetBody(entity).Post(accountingURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
