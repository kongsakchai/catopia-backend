package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kongsakchai/catopia-backend/config"
)

type DB struct {
	*sqlx.DB
}

func NewDB(cfg *config.Config) *DB {
	dbCfg := &mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Net:                  "tcp",
		Addr:                 cfg.DBHost + ":" + cfg.DBPort,
		DBName:               cfg.DBname,
		AllowNativePasswords: true,
		ParseTime:            true,
		MultiStatements:      true,
	}

	db, err := sqlx.Connect("mysql", dbCfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	return &DB{db}
}
