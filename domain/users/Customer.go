package users

type Customer struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	SurName string `json:"surName"`
	Status  string `json:"status"`
}
type Address struct {
	Id            string `json:"id"`
	UserId        string `json:"userId"`
	Address       string `json:"address"`
	AddressTypeId string `json:"addressTypeId"`
	PhoneNumber   string `json:"phoneNumber"`
}
type AddressHelper struct {
	Id          string `json:"id"`
	UserId      string `json:"userId"`
	Address     string `json:"address"`
	AddressType string `json:"addressType"`
	PhoneNumber string `json:"phoneNumber"`
}
