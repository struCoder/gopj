package model

// import (
// 	"crypto/md5"
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/go-sql-driver/mysql"
// 	"log"
// )

/*
 * isRight 此用户是否有权操作某项功能,数字越大权限越大
 * updated_at 保存用户最后一次登录的时间
**/
// func saveUser(username string, pwd string, isRight bool) bool {
// 	db, err := sql.Open("mysql", "root:123456@/gopj")
// 	defer db.Close()
// 	stmt, err := db.Prepare("insert into users(username, password, isRight, created_at, updated_at) values(?, ?, ?, ?, ?)")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer stmt.Close()

// 	if err != nil {
// 		log.Println(err)
// 		return false
// 	}
// 	// result, err := stmt.Exec(...)
// 	return false
// }

// func FindUser(email, pwd string) bool {
// 	db, err := sql.Open("mysql", "root:123456@/gopj")
// 	if err {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// 	md5Pwd = md5.Sum([]byte(pwd))
// 	rows, err := db.Query("select username, pwd from users where username=?", email)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// }
