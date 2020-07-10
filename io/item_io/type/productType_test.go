package _type

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProducts(t *testing.T) {
	resp, err := GetProductTypes()
	fmt.Println("result is: ", resp)
	assert.Nil(t, err)
}
func TestGetProductType(t *testing.T) {
	resp, err := GetProductType("T-9448a2cc-f619-48af-9a9d-6a5ad42a7faf")
	fmt.Println("result of product is: ", resp)
	assert.Nil(t, err)
}
func TestGetAllOfProductType(t *testing.T) {
	resp, err := GetAllOfProductType("T-6224368f-94c6-4d55-aaf3-38252a69a77e")
	fmt.Println("result of product is: ", resp)
	assert.Nil(t, err)
}
