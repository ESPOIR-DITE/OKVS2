package order_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOrders(t *testing.T) {
	result, err := GetOrders()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
	assert.NotNil(t, result)
}
