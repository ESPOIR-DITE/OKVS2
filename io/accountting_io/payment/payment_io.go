package payment

import (
	"OKVS2/api"
	"OKVS2/domain/accounting"
	"errors"
)

const paymentURL = api.BASE_URL + "payment"

func CreatePayment(account accounting.Payment) (accounting.Payment, error) {
	entity := accounting.Payment{}
	resp, _ := api.Rest().SetBody(account).Post(paymentURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePayment(account accounting.Payment) (accounting.Payment, error) {
	entity := accounting.Payment{}
	resp, _ := api.Rest().SetBody(account).Post(paymentURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetPayments() ([]accounting.Payment, error) {
	//entity :=Color{}
	entities := []accounting.Payment{}
	resp, _ := api.Rest().Get(paymentURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetPayment(id string) (accounting.Payment, error) {
	entity := accounting.Payment{}
	resp, _ := api.Rest().Get(paymentURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func DeletePayment(id string) (accounting.Payment, error) {
	entity := accounting.Payment{}
	resp, _ := api.Rest().Get(paymentURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
