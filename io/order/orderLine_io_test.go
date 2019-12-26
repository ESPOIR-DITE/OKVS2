package order

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOderLine(t *testing.T) {
	result, err := GetOderLine("")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
	assert.NotNil(t, result)
}
func TestGetOrderLineWithOrderId(t *testing.T) {
	result, err := GetOrderLineWithOrderId("OF-fc25885a-e70a-4c53-b76c-ca054ad55c5f")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
	assert.NotNil(t, result)
}
func TestGetOderLines(t *testing.T) {
	result, err := GetOderLines()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
	assert.NotNil(t, result)
}
