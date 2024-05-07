package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kongsakchai/catopia-backend/config"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

var database *Database

type Database struct {
	*sqlx.DB
}

func newDB() *Database {
	cfg := config.Get()

	db, err := sqlx.Connect("mysql", cfg.DBUrl+cfg.DBName+"?parseTime=true")
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

func HandlerError(err error) error {
	dbErr, ok := err.(*mysql.MySQLError)

	if ok {
		switch dbErr.Number {
		case 1062:
			return errs.NewErrorWithSkip(errs.ErrConflict, err, 2)
		}
	}

	return err
}
