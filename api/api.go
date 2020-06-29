package api

import (
	"OKVS2/config"
	"gopkg.in/resty.v1"
)

const BASE_URL string = "http://localhost:9099/OKVS/"

func Rest() *resty.Request {
	return resty.R().SetAuthToken("").
		SetHeader("Accept", "application/json").
		SetHeader("email", "email").
		SetHeader("site", "site").
		SetHeader("Content-Type", "application/json")
}

var JSON = config.ConfigWithCustomTimeFormat
