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

func UserLogin(user interface{}) (login.Login, error) {
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
