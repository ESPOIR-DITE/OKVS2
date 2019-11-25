package admin

import (
	"OKVS2/domain/users"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAdmins(t *testing.T) {
	result, err := GetAdmins()
	assert.Nil(t, err)
	fmt.Println("the result is ", result)
	assert.True(t, len(result) > 0)
}
func TestGetAdmin(t *testing.T) {
	reslt, err := GetAdmin("espoirdite@gmail.com")
	assert.Nil(t, err)
	fmt.Println("the admin is: ", reslt)
}
func TestCreateAdmin(t *testing.T) {
	//
	admin := users.Admin{"Armel@", "Armel", "tshiyoka"}
	resp, err := CreateAdmin(admin)

	assert.Nil(t, err)
	fmt.Println("new admin created: ", resp)
}
