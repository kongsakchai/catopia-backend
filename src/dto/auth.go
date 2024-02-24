package dto

type SignInDTO struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type SignUpDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Date     string `json:"date"`
}
