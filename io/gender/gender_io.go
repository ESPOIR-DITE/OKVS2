package gender

import (
	"OKVS2/api"
	"OKVS2/domain/gender"
	"errors"
)

const genderURL = api.BASE_URL + "gender"

func CreateGender(obj gender.Gender) (gender.Gender, error) {
	entity := gender.Gender{}
	resp, _ := api.Rest().SetBody(obj).Post("/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &resp)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetGender(id string) (gender.Gender, error) {
	entity := gender.Gender{}
	resp, _ := api.Rest().Get("/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &resp)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteGender(id string) (gender.Gender, error) {
	entity := gender.Gender{}
	resp, _ := api.Rest().Get("/delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &resp)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetGenders() (gender.Gender, error) {
	entity := gender.Gender{}
	resp, _ := api.Rest().Get("/reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &resp)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ResdWithGender(gende string) (gender.Gender, error) {
	entity := gender.Gender{}
	resp, _ := api.Rest().Get(genderURL + "/readWith?id=" + gende)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &resp)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
