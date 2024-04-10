package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/handler"
	"github.com/kongsakchai/catopia-backend/api/middleware"
	"github.com/kongsakchai/catopia-backend/config"
	"github.com/kongsakchai/catopia-backend/repository"
	"github.com/kongsakchai/catopia-backend/usecase"
)

type API struct {
	app *gin.Engine
}

func NewAPI() *API {
	return &API{}
}

func (api *API) Start() {
	port := config.Get().Port
	if port == "" {
		port = "8080"
	}

	app := gin.Default()
	api.app = app

	app.Use(cors.Default())
	app.Static("/images", "./upload/images")

	api.initRoute()

	app.Run(fmt.Sprintf(":%s", port))
}

func (a *API) initRoute() {

	sessionRepo := repository.NewSessionRepository()
	userRepo := repository.NewUserRepository()

	sessionUsecase := usecase.NewSessionUsecase(sessionRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	authUsecase := usecase.NewAuthUsecase(userUsecase, sessionUsecase)

	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	api := a.app.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.DELETE("/logout", middleware.AuthorizationMiddleware(sessionUsecase), authHandler.Logout)

	user := api.Group("/user", middleware.AuthorizationMiddleware(sessionUsecase))
	user.GET("/", userHandler.Get)
	user.PUT("/", userHandler.Update)
	user.PUT("/password", userHandler.UpdatePassword)
}
