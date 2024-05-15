package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/handler"
	"github.com/kongsakchai/catopia-backend/api/middleware"
	"github.com/kongsakchai/catopia-backend/config"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/repository"
	"github.com/kongsakchai/catopia-backend/usecase"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api.initRoute()

	app.Run(fmt.Sprintf(":%s", port))
}

func (a *API) initRoute() {
	db := db.GetDB()

	sessionRepo := repository.NewSessionRepository(db)
	userRepo := repository.NewUserRepository(db)
	catRepo := repository.NewCatRepository(db)
	treatmentRepo := repository.NewTreatmentRepository(db)

	fileUsecase := usecase.NewFileUsecase()
	otpUsecase := usecase.NewOTPUsecase()
	modelUsecase := usecase.NewModelUsecae()
	sessionUsecase := usecase.NewSessionUsecase(sessionRepo)
	userUsecase := usecase.NewUserUsecase(userRepo, fileUsecase, otpUsecase, modelUsecase)
	authUsecase := usecase.NewAuthUsecase(userUsecase, sessionUsecase)
	catUsecase := usecase.NewCatUsecase(catRepo, fileUsecase, modelUsecase)
	treatmentUsecase := usecase.NewTreatmentUsecase(treatmentRepo, catUsecase)
	recommendUsecase := usecase.NewRecommendUsecase(catUsecase, userUsecase)

	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase, treatmentUsecase)
	catHandler := handler.NewCatHandler(catUsecase)
	treatmentHandler := handler.NewTreatmentHandler(treatmentUsecase)
	otpHandler := handler.NewOTPHandler(otpUsecase)
	recommendHandler := handler.NewRecommendHandler(recommendUsecase)

	authMiddleware := middleware.AuthorizationMiddleware(sessionUsecase)

	api := a.app.Group("/api")
	api.GET("/healthcheck", authHandler.HealthCheck)

	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.DELETE("/logout", authMiddleware, authHandler.Logout)
	auth.GET("verify", authMiddleware, authHandler.VerifyToken)

	user := api.Group("/user")
	user.GET("", authMiddleware, userHandler.Get)
	user.PUT("", authMiddleware, userHandler.Update)
	user.POST("/answer", authMiddleware, userHandler.UserAnswer)
	user.GET("/noti", authMiddleware, userHandler.GetTreatmentNoti)

	api.PUT("/reset-password", userHandler.ResetPassword)
	api.POST("/forget-password", userHandler.ForgetPassword)
	api.POST("/otp/verify", otpHandler.VerifyOTP)

	cat := api.Group("/cat", authMiddleware)
	cat.GET("/:id", catHandler.GetByID)
	cat.GET("", catHandler.GetAll)
	cat.POST("", catHandler.Create)
	cat.PUT("/:id", catHandler.Update)
	cat.DELETE("/:id", catHandler.Delete)

	treatment := api.Group("/treatment")
	treatment.GET("/type", treatmentHandler.GetType)
	treatment.GET("/:cat_id/:id", authMiddleware, treatmentHandler.GetByID)
	treatment.GET("/:cat_id", authMiddleware, treatmentHandler.GetByCatID)
	treatment.POST("/:cat_id", authMiddleware, treatmentHandler.Create)
	treatment.PUT("/:cat_id/:id", authMiddleware, treatmentHandler.Update)
	treatment.DELETE("/:cat_id/:id", authMiddleware, treatmentHandler.Delete)

	file := api.Group("/file", authMiddleware)
	file.POST("/upload", handler.NewFileHandler().Upload)

	recommend := api.Group("/recommend")
	recommend.GET("/cat/:id", authMiddleware, recommendHandler.GetByCatID)
	recommend.GET("/cat", authMiddleware, recommendHandler.GetByUser)
}
