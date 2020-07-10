package users_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAddressTypes(t *testing.T) {
	resp, err := GetAddressTypes()
	assert.Nil(t, err)
	fmt.Println("new customer is: ", resp)
}
func TestCreateAddressType(t *testing.T) {
	//object:=AddressType{"","Home"}
	resp, err := CreateAddressType("hotel")
	assert.Nil(t, err)
	fmt.Println("new customer is: ", resp)
}
