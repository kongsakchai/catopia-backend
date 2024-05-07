package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *errs.Error `json:"error,omitempty"`
}

func NewError(c *gin.Context, err error) {
	var unwarp *errs.Error

	if errors.As(err, &unwarp) {
		c.JSON(ErrCodeToHTTPStatus[unwarp.Code], &Response{
			Success: false,
			Error:   unwarp,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, &Response{
		Success: false,
		Error:   errs.NewError(errs.ErrInternal, err),
	})
}

func New(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, &Response{
		Success: true,
		Data:    data,
	})
}
