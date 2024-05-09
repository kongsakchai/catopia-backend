package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type UpdateUser struct {
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Gender   string         `json:"gender"`
	Date     *date.JSONDate `json:"date"`
	Password string         `json:"password"`
	Profile  string         `json:"profile"`
}

type UpdatePassword struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetOTP struct {
	Username string `json:"username" binding:"required"`
}

type VerifyOTP struct {
	Code string `json:"code" binding:"required"`
	OTP  string `json:"otp" binding:"required"`
}

type UserAnswer struct {
	Answer []string `json:"answer" binding:"required"`
}
