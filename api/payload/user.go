package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type UpdateUser struct {
	Username string         `json:"username" example:"kongsakchai"`
	Email    string         `json:"email" example:"mail@mail.com"`
	Gender   string         `json:"gender" example:"male" enums:"male,female"`
	Date     *date.JSONDate `json:"date" example:"2021-01-20" format:"date"`
	Password string         `json:"password" example:"password123"`
	Profile  string         `json:"profile" example:"url of image"`
}

type UpdatePassword struct {
	Code     string `json:"code" binding:"required" example:"123456"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type GetOTP struct {
	Username string `json:"username" binding:"required" example:"kongsakchai"`
}

type VerifyOTP struct {
	Code string `json:"code" binding:"required" example:"123456"`
	OTP  string `json:"otp" binding:"required" example:"123456"`
}

type UserAnswer struct {
	Answer []float64 `json:"answer" binding:"required" example:"1,2,3,4,5"`
}
