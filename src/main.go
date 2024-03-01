package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kongsakchai/catopia-backend/src/route"
)

var db *sqlx.DB

func main() {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               "catopia",
		AllowNativePasswords: true,
	}

	print(cfg.FormatDSN())

	var err error
	db, err = sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	r := gin.Default()
	r.Use(cors.Default())

	route.RigisterRoute(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
