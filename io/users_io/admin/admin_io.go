package admin

import (
	"OKVS2/api"
	userAdmin "OKVS2/domain/users"
	"errors"
)

const adminURL = api.BASE_URL + "admin"

type Admin userAdmin.Admin

func GetAdmins() ([]Admin, error) {
	entities := []Admin{}
	resp, _ := api.Rest().Get(adminURL + "/reads")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetAdmin(id string) (Admin, error) {
	entity := Admin{}

	resp, _ := api.Rest().Get(adminURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateAdmin(entity interface{}) (bool, error) {
	resp, _ := api.Rest().SetBody(entity).Post(adminURL + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func UpdateAdmin(entity interface{}) (bool, error) {
	resp, _ := api.Rest().SetBody(entity).Post(adminURL + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func DeleteAdmin(id string) (Admin, error) {
	entity := Admin{}

	resp, _ := api.Rest().Get(adminURL + "/delete" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
