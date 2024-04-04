package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, &Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
