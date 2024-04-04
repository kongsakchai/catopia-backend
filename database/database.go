package db

import (
	"context"

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

	return &Database{db}
}

func GetDB() *Database {
	if database == nil {
		database = newDB()
	}

	return database
}

func (d *Database) UseTx(ctx context.Context, call func(*Tx) error) error {
	tx, err := d.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = call(tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
