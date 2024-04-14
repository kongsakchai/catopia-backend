package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
)

func AuthorizationMiddleware(sessionUsecase domain.SessionUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthorizationMiddleware")

		s := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")

		session, err := sessionUsecase.ValidateToken(c, token)
		if err != nil {
			response.NewErrorResponse(c, err)
			c.Abort()
			return
		}

		c.Params = append(c.Params, gin.Param{Key: "user_id", Value: strconv.FormatInt(session.UserID, 10)})
		c.Params = append(c.Params, gin.Param{Key: "session_id", Value: session.ID})
		c.Next()
	}
}
