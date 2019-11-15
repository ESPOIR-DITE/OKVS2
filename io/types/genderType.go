package types

import (
	"OKVS2/api"
	"OKVS2/domain/gender"
	"errors"
	"fmt"
)

const genderURL = api.BASE_URL + "gender"

type Gender gender.Gender

func GetGenders() ([]gender.Gender, error) {
	entities := []gender.Gender{}
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
func GetGender(id string) (gender.Gender, error) {
	entity := gender.Gender{}
	resp, _ := api.Rest().Get(braindURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func CreateGender(gender string) (Gender, error) {
	entity := Gender{}
	myType := Gender{"000", gender}
	resp, _ := api.Rest().SetBody(myType).Post(genderURL + "/create")
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
func DeleteGender(gender string) (Gender, error) {
	entity := Gender{}
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
