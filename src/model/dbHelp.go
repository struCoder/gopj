package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//get db connection

func getDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/gopj")
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	return db
}
