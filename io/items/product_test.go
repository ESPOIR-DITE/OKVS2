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
	resp, err := GetProduct("")
	fmt.Println("result of product is: ", resp)
	assert.Nil(t, err)
}
