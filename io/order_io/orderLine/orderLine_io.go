package orderLine

import (
	"OKVS2/api"
	"OKVS2/domain/orders"
	"errors"
)

const orderLineURL = api.BASE_URL + "orderline/"

func CreateOrderLine(obj orders.OrderLine) (orders.OrderLine, error) {
	entity := orders.OrderLine{}

	resp, _ := api.Rest().SetBody(obj).Post(orderLineURL + "create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetOrderLineWithOrderId(orderNumber string) ([]orders.OrderLine, error) {
	entity := []orders.OrderLine{}

	resp, _ := api.Rest().Get(orderLineURL + "readWithOrderId?id=" + orderNumber)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetOderLine(orderLineId string) (orders.OrderLine, error) {
	entity := orders.OrderLine{}

	resp, _ := api.Rest().Get(orderLineURL + "read?id=" + orderLineId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetOderLines() ([]orders.OrderLine, error) {
	entity := []orders.OrderLine{}

	resp, _ := api.Rest().Get(orderLineURL + "reads")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteOrderLine(orderLineId string) (orders.OrderLine, error) {
	entity := orders.OrderLine{}

	resp, _ := api.Rest().Get(orderLineURL + "delete?id=" + orderLineId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateOrderLine(obj orders.OrderLine) (orders.OrderLine, error) {
	entity := orders.OrderLine{}

	resp, _ := api.Rest().SetBody(obj).Post(orderLineURL + "update")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
