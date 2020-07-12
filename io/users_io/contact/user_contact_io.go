package contact

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
)

const usercontactURL = api.BASE_URL + "userContact"

func CreateUserContact(obj users.UserContact) (users.UserContact, error) {
	entity := users.UserContact{}
	resp, _ := api.Rest().SetBody(obj).Post(usercontactURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserContact(id string) (users.UserContact, error) {
	entity := users.UserContact{}
	resp, _ := api.Rest().Get(usercontactURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserContacts() ([]users.UserContact, error) {
	entity := []users.UserContact{}
	resp, _ := api.Rest().Get(usercontactURL + "/reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserContact(id string) (users.UserContact, error) {
	entity := users.UserContact{}
	resp, _ := api.Rest().Get(usercontactURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserContact(obj users.UserContact) (users.UserContact, error) {
	entity := users.UserContact{}
	resp, _ := api.Rest().SetBody(obj).Post(usercontactURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
