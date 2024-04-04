package response

import (
	"errors"

	"github.com/gin-gonic/gin"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

func NewErrorResponse(c *gin.Context, err error) {
	var unwarp *errs.Error

	if errors.As(err, &unwarp) {
		c.JSON(unwarp.Code, &Response{
			Code:    unwarp.Code,
			Message: unwarp.Message,
		})
		return
	}

	c.AbortWithStatusJSON(errs.ErrInternal, &Response{
		Code:    errs.ErrInternal,
		Message: "Internal Server Error",
	})
}
