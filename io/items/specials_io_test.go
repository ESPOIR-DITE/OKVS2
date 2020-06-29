package items

import (
	"OKVS2/domain/items"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSpecial(t *testing.T) {
	obje := items.Specials{"", "", "", "", "", "", "", 0}
	result, err := CreateSpecial(obje)
	fmt.Println("result ", result)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
