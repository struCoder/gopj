package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

/*
 * isRight 此用户是否有权操作某项功能,数字越大权限越大
 * updated_at 保存用户最后一次登录的时间
**/

func saveUser(username string, pwd string, isRight int) bool {
	db := getDb()
	defer db.Close()
	stmt, err := db.Prepare("insert into users(username, password, isRight, created_at, updated_at) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return false
	}
	nowTime := time.Now()
	stmt.Exec(username, pwd, 3, nowTime, nowTime)
	return true
}

func FindUser(email, md5Pwd string) (bool, map[string]UserInfo) {
	var _id, _username, _pwd string
	db := getDb()
	defer db.Close()
	_err := db.QueryRow("select id, username, password from users where username=?", email).Scan(&_id, &_username, &_pwd)

	switch {
	case _err == sql.ErrNoRows:
		log.Printf("no user")
	case _err != nil:
		log.Fatal(_err)
	default:
		if md5Pwd == _pwd {
			userInfoStruct := make(map[string]UserInfo)
			userInfoStruct["user"] = UserInfo{_id, _username}
			return true, userInfoStruct
		}
	}

	return false, nil
}
