package users

type LoginHelper struct {
	Email   string `json:"email"`
	Pasword string `json:"pasword"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}
