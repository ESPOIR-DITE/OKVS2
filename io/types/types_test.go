package types

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
	resp, err := CreateType("shoes")
	assert.Nil(t, err)
	fmt.Println("result is: ", resp)
}
