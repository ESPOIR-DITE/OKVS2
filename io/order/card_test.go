package order

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCard(t *testing.T) {
	resp, err := GetCard("CF-5dff01ed-665f-4710-bac0-0721cb154281")
	fmt.Println("result of Card is: ", resp)
	assert.Nil(t, err)

}
func TestGetCustomer(t *testing.T) {
	resp, err := GetCardWithCustId("espoir@dite.com")
	fmt.Println("result of Card is: ", resp)
	assert.Nil(t, err)
}
