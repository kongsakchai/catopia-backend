package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kongsakchai/catopia-backend/src/handler"
	"github.com/kongsakchai/catopia-backend/src/repository"
	"github.com/kongsakchai/catopia-backend/src/usecase"
)

func RigisterRoute(r *gin.Engine, db *sqlx.DB) {
	api := r.Group("/api")

	userRepo := repository.NewUserRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo)

	authHandler := handler.NewAuthHandler(authUsecase)

	AuthRoute(api, authHandler)
}

func AuthRoute(r *gin.RouterGroup, controller *handler.AuthHandler) {
	api := r.Group("/auth")

	api.POST("/sign-up", controller.SignUp)
	api.POST("/sign-in", controller.SignIn)
}
