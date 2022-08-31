package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/majoo?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
