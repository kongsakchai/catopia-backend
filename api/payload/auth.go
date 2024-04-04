package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Regis struct {
	Username string         `json:"username" binding:"required"`
	Password string         `json:"password" binding:"required"`
	Email    string         `json:"email" binding:"required"`
	Gender   string         `json:"gender" binding:"required"`
	Date     *date.JSONDate `json:"date" binding:"required"`
}
