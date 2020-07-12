package users_io

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
	"fmt"
)

const roleURL = api.BASE_URL + "role"

func GetRoles() ([]users.Roles, error) {
	entities := []users.Roles{}
	resp, _ := api.Rest().Get(roleURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetRole(id string) (users.Roles, error) {
	entity := users.Roles{}
	resp, _ := api.Rest().Get(roleURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func CreateRole(gender users.Roles) (users.Roles, error) {
	entity := users.Roles{}
	resp, _ := api.Rest().SetBody(gender).Post(roleURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	fmt.Println(" we are creating Gender", err)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateRole(gender string) (users.Roles, error) {
	entity := users.Roles{}
	resp, _ := api.Rest().SetBody(gender).Post(roleURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	fmt.Println(" we are creating Gender", err)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteRole(gender string) (users.Roles, error) {
	entity := users.Roles{}
	resp, _ := api.Rest().Get(roleURL + "/delete?id=" + gender)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	fmt.Println(" we are creating Gender", err)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
