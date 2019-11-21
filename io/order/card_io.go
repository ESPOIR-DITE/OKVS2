package order

import (
	"OKVS2/api"
	"OKVS2/domain/orders"
	"errors"
)

const cardURL = api.BASE_URL + "card"

func CreateCard(myEntity orders.Card) (orders.Card, error) {
	entity1 := orders.Card{"00", myEntity.ItemId, myEntity.CustomerId}
	entity := orders.Card{}
	resp, _ := api.Rest().SetBody(entity1).Post(cardURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetCardWithItemId(itemId string) (orders.Card, error) {
	entity := orders.Card{}
	resp, _ := api.Rest().Get(cardURL + "/readWithItemId?id=" + itemId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetCardWithCustId(itemId string) ([]orders.Card, error) {
	entity := []orders.Card{}
	resp, _ := api.Rest().Get(cardURL + "/readWithCustId?id=" + itemId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetCard(id string) ([]orders.Card, error) {
	entity := []orders.Card{}
	resp, _ := api.Rest().Get(cardURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetCards() ([]orders.Card, error) {
	entity := []orders.Card{}
	resp, _ := api.Rest().Get(cardURL + "/reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCard(id string) (orders.Card, error) {
	entity := orders.Card{}
	resp, _ := api.Rest().Get(cardURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCard(myEntity interface{}) (orders.Card, error) {
	entity := orders.Card{}
	resp, _ := api.Rest().SetBody(myEntity).Post(cardURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
