package _type

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTypes(t *testing.T) {
	resp, err := GetTypes()
	fmt.Println("result is: ", resp)
	assert.Nil(t, err)
}
func TestCreateType(t *testing.T) {
	resp, err := CreateType("pants tshirt")
	assert.Nil(t, err)
	fmt.Println("result is: ", resp)
}
