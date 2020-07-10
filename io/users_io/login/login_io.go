package login

import (
	"OKVS2/api"
	"OKVS2/domain/users"
	"errors"
	"fmt"
)

const loginURL = api.BASE_URL + "login"

type userLogin users.Login
type userDetail users.LoginHelper

func UserLogin(user users.Login) (users.Login, error) {
	fmt.Println("user: ", user)
	entity := users.Login{}
	resp, _ := api.Rest().SetBody(user).Post(loginURL + "/log")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserEmail(pasword string) (users.Login, error) {
	entity := users.Login{}
	resp, _ := api.Rest().Get(loginURL + "/readwithpassword?id=" + pasword)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserWithEmail(email string) (users.Login, error) {
	entity := users.Login{}
	resp, _ := api.Rest().Get(loginURL + "/read?id=" + email)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateLogin(user users.Login) (users.Login, error) {
	//fmt.Println("user: ", user)
	entity := users.Login{}

	resp, _ := api.Rest().SetBody(user).Post(loginURL + "/update")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UniversalLogin(user users.Login) (users.Login, error) {
	entity := users.Login{}
	resp, _ := api.Rest().SetBody(user).Post(loginURL + "/univelogin")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
