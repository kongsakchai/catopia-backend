package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type UpdateUser struct {
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Gender   string         `json:"gender"`
	Date     *date.JSONDate `json:"date"`
	Password string         `json:"password"`
}

type UpdatePassword struct {
	Password string `json:"password" binding:"required"`
}
