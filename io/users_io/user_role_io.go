package users_io

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
	"fmt"
)

const UserroleURL = api.BASE_URL + "user_role"

func GetUserRoles() ([]users.UserRole, error) {
	entities := []users.UserRole{}
	resp, _ := api.Rest().Get(UserroleURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetUserRole(id string) (users.UserRole, error) {
	entity := users.UserRole{}
	resp, _ := api.Rest().Get(UserroleURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateUserRole(gender users.UserRole) (users.UserRole, error) {
	entity := users.UserRole{}
	resp, _ := api.Rest().SetBody(gender).Post(UserroleURL + "/create")
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
func UpdateUserRole(gender string) (users.UserRole, error) {
	entity := users.UserRole{}
	resp, _ := api.Rest().SetBody(gender).Post(UserroleURL + "/update")
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
func DeleteUserRole(gender string) (users.UserRole, error) {
	entity := users.UserRole{}
	resp, _ := api.Rest().Get(UserroleURL + "/delete?id=" + gender)
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
