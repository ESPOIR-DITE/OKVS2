package joins

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOneItemDetails(t *testing.T) {
	result, err := GetOneItemDetails("PF-01eb88e7-3396-4eac-a235-5960608a9579")
	assert.Nil(t, err)
	fmt.Println("new customer is: ", result)
}
