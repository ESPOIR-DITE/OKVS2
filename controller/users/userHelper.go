package users

import (
	"OKVS2/domain/users"
	"OKVS2/io/order_io/card"
	"OKVS2/io/users_io"
	"OKVS2/io/users_io/admin"
	"OKVS2/io/users_io/login"
	"fmt"
)

type CardeData struct {
	Mesage string
	Class  string
}

func GetUserDetails(email string) (string, string, bool, users.Customer) {

	var message string
	var class string
	var Manager = false
	var user users.Customer

	userLog, err := login.GetUserWithEmail(email)
	if userLog.UserType == "customer" {
		user, err = users_io.GetCustomer(email)
		if err != nil {
			fmt.Println("error reading GetCustomer in if userLog.UserType==customer")
		}
	} else if userLog.UserType == "admin" {
		_, err := admin.GetAdmin(email)
		if err != nil {
			fmt.Println("error reading GetCustomer in if userLog.UserType==customer2")
		} else {
			Manager = true
		}
	}

	cardDetails, err := card.GetCardWithCustId(email)
	fmt.Println("User card>>: ", cardDetails)
	if err != nil {
		fmt.Println("User may not have logIn or may not have ordered yet ")
	}
	var itemIdfromcard string
	for _, valeu := range cardDetails {
		itemIdfromcard = valeu.ItemId
	}
	if itemIdfromcard != "" {
		//app.ErrorLog.Println(err.Error())
		message = "You have something in your Card please click on the trolley icon to view your card"
		class = "primary"
	}

	return message, class, Manager, user
}
