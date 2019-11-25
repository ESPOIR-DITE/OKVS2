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
	resp, err := GetProduct("PF-3f390a44-3e43-4a03-959a-585118bbbeac")
	fmt.Println("result of product is: ", resp)
	assert.Nil(t, err)
}
