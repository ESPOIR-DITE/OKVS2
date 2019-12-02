package users

type Customer struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	SurName string `json:"surName"`
	Status  string `json:"status"`
}
type Address struct {
	Id          string `json:"id"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}
