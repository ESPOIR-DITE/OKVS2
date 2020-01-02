package login

import (
	"OKVS2/api"
	"OKVS2/domain/login"
	"errors"
	"fmt"
)

const loginURL = api.BASE_URL + "login"

type userLogin login.Login
type userDetail login.LoginHelper

func UserLogin(user login.Login) (login.Login, error) {
	fmt.Println("user: ", user)
	entity := login.Login{}
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
func GetUserEmail(pasword string) (login.Login, error) {
	entity := login.Login{}
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
func GetUserWithEmail(email string) (login.Login, error) {
	entity := login.Login{}
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
func UpdateLogin(user login.Login) (login.Login, error) {
	//fmt.Println("user: ", user)
	entity := login.Login{}

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
func UniversalLogin(email, password string) (login.Login, error) {
	entity := login.Login{}
	resp, _ := api.Rest().Get(loginURL + "/univelogin?email=" + email + "&password=" + password)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
