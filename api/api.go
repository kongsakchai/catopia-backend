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

	cors := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	})
	app.Use(cors)

	// app.Use(middleware.CORSMiddleware())

	app.Static("/images", "./uploads")

	api.initRoute()

	app.Run(fmt.Sprintf(":%s", port))
}

func (a *API) initRoute() {

	sessionRepo := repository.NewSessionRepository()
	userRepo := repository.NewUserRepository()
	catRepo := repository.NewCatRepository()
	treatmentRepo := repository.NewTreatmentRepository()

	sessionUsecase := usecase.NewSessionUsecase(sessionRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	authUsecase := usecase.NewAuthUsecase(userUsecase, sessionUsecase)
	catUsecase := usecase.NewCatUsecase(catRepo)
	treatmentUsecase := usecase.NewTreatmentUsecase(treatmentRepo, catUsecase)

	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	catHandler := handler.NewCatHandler(catUsecase)
	treatmentHandler := handler.NewTreatmentHandler(treatmentUsecase)

	authMiddleware := middleware.AuthorizationMiddleware(sessionUsecase)

	api := a.app.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.DELETE("/logout", authMiddleware, authHandler.Logout)

	user := api.Group("/user", authMiddleware)
	user.GET("", userHandler.Get)
	user.PUT("", userHandler.Update)
	user.PUT("/password", userHandler.UpdatePassword)

	cat := api.Group("/cat", authMiddleware)
	cat.GET("/:id", catHandler.GetByID)
	cat.GET("", catHandler.GetAll)
	cat.POST("", catHandler.Create)
	cat.PUT("/:id", catHandler.Update)
	cat.DELETE("/:id", catHandler.Delete)

	treatment := api.Group("/treatment", authMiddleware)
	treatment.GET("/:cat_id/:id", treatmentHandler.GetByID)
	treatment.GET("/:cat_id", treatmentHandler.GetByCatID)
	treatment.POST("/:cat_id", treatmentHandler.Create)
	treatment.PUT("/:cat_id/:id", treatmentHandler.Update)
	treatment.DELETE("/:cat_id/:id", treatmentHandler.Delete)

	file := api.Group("/file", authMiddleware)
	file.POST("/upload", handler.NewFileHandler().Upload)
}
