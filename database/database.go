package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kongsakchai/catopia-backend/config"
)

var database *Database

type Database struct {
	*sqlx.DB
}

type Tx = sqlx.Tx

func newDB() *Database {
	cfg := config.Get()

	dbCfg := &mysql.Config{
		User:      cfg.DBUser,
		Passwd:    cfg.DBPassword,
		Net:       "tcp",
		Addr:      cfg.DBHost + ":" + cfg.DBPort,
		DBName:    cfg.DBname,
		ParseTime: true,
	}

	db, err := sqlx.Connect("mysql", dbCfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	return &Database{db}
}

func GetDB() *Database {
	if database == nil {
		database = newDB()
	}

	return database
}
