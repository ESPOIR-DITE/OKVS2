package gender

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
)

const itemGenderURL = api.BASE_URL + "itemGender"

type ItemGender items.ItemGender

func GetItemGenders() ([]items.ItemGender, error) {
	entities := []items.ItemGender{}
	resp, _ := api.Rest().Get(itemGenderURL + "/reads")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetItemGender(id string) (items.ItemGender, error) {
	entity := items.ItemGender{}
	resp, _ := api.Rest().Get(itemGenderURL + "/read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	//fmt.Println(" we are about to reading ItemGender <<int item gender>>", entity)
	return entity, nil

}
func CreateItemGender(itemGender items.ItemColor) (items.ItemGender, error) {
	entity := items.ItemGender{}
	resp, _ := api.Rest().SetBody(itemGender).Post(itemGenderURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())

	}
	return entity, nil
}
func DeleteItemGender(braind string) (items.ItemGender, error) {
	entity := items.ItemGender{}
	resp, _ := api.Rest().Get(itemGenderURL + "/delete?id=" + braind)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadItemGenderWithItemId(itemId string) (items.ItemGender, error) {
	entity := items.ItemGender{}
	resp, _ := api.Rest().Get(itemGenderURL + "/readWithItemId?itemId=" + itemId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateItemGender(itemGender items.ItemGender) (items.ItemGender, error) {
	entity := items.ItemGender{}
	resp, _ := api.Rest().SetBody(itemGender).Post(itemGenderURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
