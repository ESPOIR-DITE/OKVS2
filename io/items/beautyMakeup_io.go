package items

import (
	"OKVS2/api"
	"OKVS2/domain/items"
	"errors"
	"fmt"
)

const beautyURL = api.BASE_URL + "myitems"

type BeautyItem items.BeautyMakeup
type BeutyHleper items.MyItemHelper

func CreatBeatyHelper(bh interface{}) (bool, error) {
	//fmt.Println(bh)
	resp, _ := api.Rest().SetBody(bh).Post(beautyURL + "/creatwithfile")
	//fmt.Println("result: ", resp)
	if resp.IsError() {
		fmt.Println("checking if there any error: ", resp)
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func GetBeauties() ([]BeautyItem, error) {
	entities := []BeautyItem{}
	resp, _ := api.Rest().Get(beautyURL + "/reads")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)

	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
func GetBeaty(id string) (BeautyItem, error) {
	entity := BeautyItem{}

	resp, _ := api.Rest().Get(beautyURL + "/read" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateBeauty(beaute interface{}) (BeautyItem, error) {
	entity := BeautyItem{}

	resp, _ := api.Rest().SetBody(beaute).Post(beautyURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBeauty(id string) (BeautyItem, error) {
	entity := BeautyItem{}
	resp, _ := api.Rest().Get(beautyURL + "delete" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateBeauty(beauty interface{}) (BeautyItem, error) {
	entity := BeautyItem{}
	resp, _ := api.Rest().SetBody(beauty).Post(beautyURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
