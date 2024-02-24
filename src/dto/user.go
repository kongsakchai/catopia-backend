package dto

type UserDTO struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Name     string `json:"name" db:"name"`
	Gender   string `json:"gender" db:"gender"`
	Date     string `json:"date" db:"date"`
}
