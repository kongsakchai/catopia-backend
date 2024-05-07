package main

import (
	_ "os"

	"github.com/kongsakchai/catopia-backend/api"
	"github.com/kongsakchai/catopia-backend/config"
	db "github.com/kongsakchai/catopia-backend/database"
	ort "github.com/yalue/onnxruntime_go"
)

func main() {
	cfg := config.Get()

	ort.SetSharedLibraryPath(cfg.ONNXPath)
	err := ort.InitializeEnvironment()
	if err != nil {
		panic(err)
	}
	defer ort.DestroyEnvironment()

	database := db.GetDB()
	defer database.Close()

	server := api.NewAPI()
	server.Start()
}
