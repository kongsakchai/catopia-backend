package dto

type Response struct {
	Code    int    `json:"code"`
	Sucess  bool   `json:"success"`
	Result  any    `json:"result"`
	Message string `json:"message"`
}
