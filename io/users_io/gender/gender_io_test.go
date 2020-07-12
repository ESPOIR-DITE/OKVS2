package gender

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGender(t *testing.T) {
	resp, err := GetGender("espoirditekemena@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}
func TestGetGenders(t *testing.T) {
	resp, err := GetGenders()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}
