package login

import (
	"OKVS2/domain/login"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserLogin(t *testing.T) {
	result := login.LoginHelper{"espoirdite@gmail.com", "LF-4d728550-8ec9-40a6-ace4-1bca27128b52"}
	rspo, err := UserLogin(result)
	assert.Nil(t, err)
	fmt.Println("result is: ", rspo)
}
