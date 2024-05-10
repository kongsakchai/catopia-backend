package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kongsakchai/catopia-backend/config"
)

func main() {
	log.Println("Migration Start")
	cfg := config.Get()

	db, err := sql.Open("mysql", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.New("file://"+cfg.MigrationPath, "mysql://"+cfg.DBUrl+cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	log.Println("Migration done")
}
