package makeUp

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const itemviewURL = api.BASE_URL + "join"

func GetAllItems() ([]items.ItemView, error) {
	entities := []items.ItemView{}
	resp, _ := api.Rest().Get(itemviewURL + "/view")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
