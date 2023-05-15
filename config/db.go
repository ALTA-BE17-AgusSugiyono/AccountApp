package config

import (
	"database/sql"
)

var DB *sql.DB

func Init() error {
	var err error
	DB, err = InitDB()
	if err != nil {
		return err
	}
	return nil
}
