package users

import (
	"OKVS2/domain/users"
	"OKVS2/io/login"
	"OKVS2/io/order"
	"OKVS2/io/users_io/admin"
	"OKVS2/io/users_io/customer"
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
	if userLog.UserTupe == "customer" {
		user, err = customer.GetCustomer(email)
		if err != nil {
			fmt.Println("error reading GetCustomer in if userLog.UserTupe==customer")
		}
	} else if userLog.UserTupe == "admin" {
		_, err := admin.GetAdmin(email)
		if err != nil {
			fmt.Println("error reading GetCustomer in if userLog.UserTupe==customer2")
		} else {
			Manager = true
		}
	}

	cardDetails, err := order.GetCardWithCustId(email)
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
