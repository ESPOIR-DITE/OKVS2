package order

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOrderStatues(t *testing.T) {
	result, err := GetOrderStatues()
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
	assert.NotNil(t, result)
}
func TestGetOrderStatus(t *testing.T) {
	result, err := GetOrderStatus("OF-fc25885a-e70a-4c53-b76c-ca054ad55c5f")
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
	assert.NotNil(t, result)
}
func TestGetWithOrderId(t *testing.T) {
	result, err := GetWithOrderId("OF-fc25885a-e70a-4c53-b76c-ca054ad55c5f")
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
	assert.NotNil(t, result)
}
