package order

import (
	"OKVS2/api"
	"OKVS2/domain/orders"
	"errors"
)

const orderStatusURL = api.BASE_URL + "orderstatus/"

func CreateOrderStatus(obj orders.OrderStatus) (orders.OrderStatus, error) {
	entity := orders.OrderStatus{}
	resp, _ := api.Rest().SetBody(obj).Post(orderStatusURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetWithOrderId(orderIs string) (orders.OrderStatus, error) {
	entity := orders.OrderStatus{}
	resp, _ := api.Rest().Get(orderStatusURL + "readWithOrderId?id=" + orderIs)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetOrderStatus(orderStarusIs string) (orders.OrderStatus, error) {
	entity := orders.OrderStatus{}
	resp, _ := api.Rest().Get(orderStatusURL + "read?id=" + orderStarusIs)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetOrderStatues() ([]orders.OrderStatus, error) {
	entity := []orders.OrderStatus{}
	resp, _ := api.Rest().Get(orderStatusURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteOrderStatus(orderIs string) (orders.OrderStatus, error) {
	entity := orders.OrderStatus{}
	resp, _ := api.Rest().Get(orderStatusURL + "delete?id=" + orderIs)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateOrderStatus(obj orders.OrderStatus) (orders.OrderStatus, error) {
	entity := orders.OrderStatus{}
	resp, _ := api.Rest().SetBody(obj).Post(orderStatusURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
