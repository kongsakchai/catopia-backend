package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type UpdateUser struct {
	Username string         `json:"username" binding:"required"`
	Email    string         `json:"email" binding:"required"`
	Gender   string         `json:"gender" binding:"required"`
	Date     *date.JSONDate `json:"date" binding:"required"`
}

type UpdatePassword struct {
	Password string `json:"password" binding:"required"`
}
