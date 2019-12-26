package customer

import (
	"OKVS2/domain/gender"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCustomerGender(t *testing.T) {
	customerGender := gender.CustomerGender{"ooss", "ueueue", "29"}
	resp, err := CreateCustomerGender(customerGender)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}
func TestUpdateCustomerGender(t *testing.T) {
	customerGender := gender.CustomerGender{"ooss", "000122222", "29"}
	resp, err := UpdateCustomerGender(customerGender)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}
func TestDeleteCustomerGender(t *testing.T) {
	//customerGender:=gender.CustomerGender{"ooss","ueueue","29"}
	resp, err := DeleteCustomerGender("ooss")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}
func TestGetCustomerGenders(t *testing.T) {
	resp, err := GetCustomerGenders()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}
