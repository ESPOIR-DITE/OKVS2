package gender

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
	"fmt"
)

const genderURL = api.BASE_URL + "gender"

func GetGenders() ([]users.Gender, error) {
	entities := []users.Gender{}
	resp, _ := api.Rest().Get(genderURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetGender(id string) (users.Gender, error) {
	entity := users.Gender{}
	resp, _ := api.Rest().Get(genderURL + "/read?id=" + id)
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
func CreateGender(gender string) (users.Gender, error) {
	entity := users.Gender{}
	resp, _ := api.Rest().SetBody(gender).Post(genderURL + "/create")
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
func DeleteGender(gender string) (users.Gender, error) {
	entity := users.Gender{}
	resp, _ := api.Rest().Get(genderURL + "/delete?id=" + gender)
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
