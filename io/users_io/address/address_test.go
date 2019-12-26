package users_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAddress(t *testing.T) {
	resp, err := GetAddress("espoirditekemena@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}
