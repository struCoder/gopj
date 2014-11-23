package model

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
 * param inputName: 投诉人
 * param inputTel: 投诉人电话
 * param beInputName: 被投诉人或群体
 * param beComplaintHome: 被投诉人的地址
 * param inputReason: 投诉的理由
 * param dealPerson: 处理人
 */

//here, we use db.Prepare(query) to execute some sql querys, it can prevent SQL injection
func SaveComplaint(inputName, inputTel, beInputName, beComplaintHome, inputReason, dealPerson string) (bool, string) {
	db := getDb()
	defer db.Close()
	query := "INSERT INTO complaint (name, phone, be_complainted, address, reason, deal_person)" +
		"VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return false, "sql errors"
	}
	defer stmt.Close()
	stmt.Exec(inputName, inputTel, beInputName, beComplaintHome, inputReason, dealPerson)
	return true, ""
}
