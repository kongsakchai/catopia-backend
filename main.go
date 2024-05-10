package main

import (
	_ "os"

	_ "github.com/kongsakchai/catopia-backend/docs"

	"github.com/kongsakchai/catopia-backend/api"
	"github.com/kongsakchai/catopia-backend/config"
	db "github.com/kongsakchai/catopia-backend/database"
	ort "github.com/yalue/onnxruntime_go"
)

// @title Catopia API
// @version 1.0
// @description This is a Catopia API of CPE Senior Project.
// @termsOfService http://somewhere.com/

// @contact.name CPE34 - Catopia

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
