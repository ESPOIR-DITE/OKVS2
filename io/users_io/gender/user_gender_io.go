package gender

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
	"fmt"
)

const usergenderURL = api.BASE_URL + "user_gender"

func GetUserGenders() ([]users.Gender, error) {
	entities := []users.Gender{}
	resp, _ := api.Rest().Get(usergenderURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetUserGender(id string) (users.Gender, error) {
	entity := users.Gender{}
	resp, _ := api.Rest().Get(usergenderURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	fmt.Println(" we are reading Gender <<in Gender IO>>", entity)
	return entity, nil

}
func CreateUserGender(gender string) (users.Gender, error) {
	entity := users.Gender{}
	resp, _ := api.Rest().SetBody(gender).Post(usergenderURL + "/create")
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
func UpdateUserGender(gender string) (users.Gender, error) {
	entity := users.Gender{}
	resp, _ := api.Rest().SetBody(gender).Post(usergenderURL + "/update")
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
func DeleteUserGender(gender string) (users.Gender, error) {
	entity := users.Gender{}
	resp, _ := api.Rest().Get(usergenderURL + "/delete?id=" + gender)
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
