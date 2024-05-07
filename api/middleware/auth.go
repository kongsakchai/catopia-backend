package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

func AuthorizationMiddleware(sessionUsecase domain.SessionUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthorizationMiddleware")

		s := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")

		if token == "" {
			response.NewError(c, errs.NewError(errs.ErrUnauthorized, fmt.Errorf("Unauthorized")))
			c.Abort()
			return
		}

		session, err := sessionUsecase.ValidateToken(c, token)
		if err != nil {
			response.NewError(c, err)
			c.Abort()
			return
		}

		c.Params = append(c.Params, gin.Param{Key: "user_id", Value: strconv.FormatInt(session.UserID, 10)})
		c.Params = append(c.Params, gin.Param{Key: "session_id", Value: session.ID})
		c.Next()
	}
}
