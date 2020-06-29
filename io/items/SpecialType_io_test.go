package items

import (
	"OKVS2/domain/items"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSpecialType(t *testing.T) {
	obj := items.SpecialType{"", "Saisson selle", ""}
	result, err := CreateSpecialType(obj)
	fmt.Println("result ", result)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
func TestGetSpecialTypes(t *testing.T) {
	result, err := GetSpecialTypes()
	fmt.Println("result ", result)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
