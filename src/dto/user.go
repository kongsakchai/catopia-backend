package dto

type SignInDTO struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type SignUpDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name" db:"name"`
	Gender   string `json:"gender" db:"gender"`
	Date     string `json:"date" db:"date"`
}

type UserDTO struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Name     string `json:"name" db:"name"`
	Gender   string `json:"gender" db:"gender"`
	Date     string `json:"date" db:"date"`
}
