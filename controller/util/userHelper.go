package util

import (
	"OKVS2/io/users_io/admin"
)

func VerifyAdmin(email string) bool {
	admin, err := admin.GetAdmin(email)
	if err != nil {
		return false
	} else if admin.Email != "" {
		return true
	}
	return false
}
