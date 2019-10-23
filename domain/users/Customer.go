package users

type Customer struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	SurName string `json:"surName"`
	Status  string `json:"status"`
}
