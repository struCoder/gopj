package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
 * isRight 此用户是否有权操作某项功能
 * updated_at 保存用户最后一次登录的时间
**/
func saveUser(username string, pwd string, isRight bool) bool {
	stmt, err := db.Prepare("insert into users(username, password, isRight, created_at, updated_at) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}
	// result, err := stmt.Exec(...)
}

func findUser(email, pwd string) bool {
	err := db.QueryRow("select username, pwd from users where email=?", email).Scan(&email)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no ")
		return false
	case err != nil:
		log.Fatal("fuck")
	}
	return true
}
