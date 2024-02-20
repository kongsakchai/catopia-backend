package model

type UserModel struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Name     string `json:"name" db:"name"`
	Salt     string `json:"salt" db:"salt"`
	Gender   string `json:"gender" db:"gender"`
	Date     string `json:"date" db:"date"`
}
