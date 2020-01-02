package order

import (
	"OKVS2/domain/orders"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateStatus(t *testing.T) {
	obj := orders.Status{"", "Completed"}
	result, err := CreateStatus(obj)
	assert.Nil(t, err)
	fmt.Println("result: ", result)
	assert.NotNil(t, result)
}
func TestGetStatues(t *testing.T) {
	result, err := GetStatues()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
	assert.NotNil(t, result)
}
func TestDeleteOrderStatus(t *testing.T) {

}
