package items

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProducts(t *testing.T) {
	resp, err := GetProducts()
	fmt.Println("result of product is: ", resp)
	assert.Nil(t, err)
}
func TestGetProduct(t *testing.T) {
	resp, err := GetProduct("PF-327101f4-1dc5-42f0-937e-f0e4ebf8630d")
	fmt.Println("result of product is: ", resp)
	assert.Nil(t, err)
}
