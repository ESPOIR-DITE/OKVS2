package size

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPtoductSizeWithItemId(t *testing.T) {
	result, err := GetPtoductSizeWithItemId("PF-01eb88e7-3396-4eac-a235-5960608a9579")
	fmt.Println("result is: ", result)
	assert.Nil(t, err)
}
