package image

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadWithImageId(t *testing.T) {
	result, err := ReadWithImageId("IF-659bc510-69c4-4d2a-b4f8-a20b55884190")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
