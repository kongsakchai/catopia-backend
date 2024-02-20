package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kongsakchai/catopia-backend/src/controller"
	"github.com/kongsakchai/catopia-backend/src/repository"
	"github.com/kongsakchai/catopia-backend/src/usecase"
)

var db *sqlx.DB

func main() {
	cfg := mysql.Config{
		User:   "root",     // os
		Passwd: "12345678", // os
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "catopia",
	}

	var err error
	db, err = sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	r := gin.Default()
	r.Use(cors.Default())
	routeRigister(r)

	r.Run()
}

func routeRigister(r *gin.Engine) {
	api := r.Group("/api")

	user(api)

}

func user(r *gin.RouterGroup) {
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	controller := controller.NewUserController(usecase)

	r.POST("/sign-up", controller.SignUp)
	r.POST("/sign-in", controller.SignIn)
}
