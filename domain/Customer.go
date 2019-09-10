package domain

type Customer struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	surname string `json:"surname"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Status  string `json:"status"`
}
