package customer

import (
	"OKVS2/domain/users"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	cust := users.Customer{"Armel@", "tshiyo", "kalos", "active"}
	resp, err := CreateCustomer(cust)
	assert.Nil(t, err)
	fmt.Println("new customer is: ", resp)
}
func TestGetCustomer(t *testing.T) {
	resp, err := GetCustomer("espoirditekemena@gmail.com")
	assert.Nil(t, err)
	fmt.Println("new customer is: ", resp)
}
func TestGetCustomers(t *testing.T) {
	resp, err := GetCustomers()
	assert.Nil(t, err)
	fmt.Println(resp)
}
