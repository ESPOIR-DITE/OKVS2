package contact

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
)

const contacttypeURL = api.BASE_URL + "contactType"

func CreateContactType(obj users.ContactType) (users.ContactType, error) {
	entity := users.ContactType{}
	resp, _ := api.Rest().SetBody(obj).Post(contacttypeURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetContactType(id string) (users.ContactType, error) {
	entity := users.ContactType{}
	resp, _ := api.Rest().Get(contacttypeURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetContactTypes() ([]users.ContactType, error) {
	entity := []users.ContactType{}
	resp, _ := api.Rest().Get(contacttypeURL + "/reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteContactType(id string) (users.ContactType, error) {
	entity := users.ContactType{}
	resp, _ := api.Rest().Get(contacttypeURL + "/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateContactType(obj users.ContactType) (users.ContactType, error) {
	entity := users.ContactType{}
	resp, _ := api.Rest().SetBody(obj).Post(contacttypeURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
