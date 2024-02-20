package dto

type Response struct {
	Sucess  bool   `json:"success"`
	Result  any    `json:"result"`
	Message string `json:"message"`
}
