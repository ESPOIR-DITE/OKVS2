package users

import "time"

type AddressType struct {
	Id          string `json:"id"`
	AddressType string `json:"addressType"`
}
type UserAddress struct {
	Id          string `json:"id"`
	UserId      string `json:"userId"`
	Address     string `json:"address"`
	AddressType string `json:"addressType"`
	PhoneNumber string `json:"phoneNumber"`
}

type ContactType struct {
	Id      string `json:"id"`
	Contact string `json:"contact"`
}
type UserContact struct {
	Id          string `json:"id"`
	ContactType string `json:"contactType"`
	Contact     string `json:"contact"`
}

type UserGender struct {
	Id       string `json:"id"`
	UserId   string `json:"userId"`
	GenderId string `json:"genderId"`
	Age      string `json:"age"`
}
type Gender struct {
	GenderId   string `json:"genderId"`
	GenderName string `json:"genderName"`
}

type User struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	SurName string `json:"surName"`
	Status  string `json:"status"`
}

type Roles struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type UserRole struct {
	Id          string    `json:"id"`
	Email       string    `json:"email"`
	RoleId      string    `json:"roleId"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Actor       string    `json:"actor"`
}
