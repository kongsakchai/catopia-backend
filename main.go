package main

import (
	"github.com/kongsakchai/catopia-backend/api"
	db "github.com/kongsakchai/catopia-backend/database"
)

func main() {
	database := db.GetDB()
	defer database.Close()

	server := api.NewAPI()
	server.Start()
}
