package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type Login struct {
	Username string `json:"username" binding:"required" example:"kongsakchai"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Regis struct {
	Username string         `json:"username" binding:"required" example:"kongsakchai"`
	Password string         `json:"password" binding:"required" example:"password123"`
	Email    string         `json:"email" binding:"required" example:"email@email.com"`
	Gender   string         `json:"gender" binding:"required" example:"male" enums:"male,female"`
	Date     *date.JSONDate `json:"date" binding:"required" example:"2021-01-20" format:"date"`
}
