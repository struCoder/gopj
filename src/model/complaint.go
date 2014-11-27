package model

import (
	"database/sql"
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

func GetMsgByLimit(start, end int) (map[int]ComplaintMsg, int) {
	db := getDb()
	defer db.Close()
	i := 1
	var Num int
	db.QueryRow("select count(*) from complaint").Scan(&Num)
	rows, err := db.Query("select id, name, phone, be_complainted, reason, status, deal_person from complaint limit ?, ?", start, end)
	defer rows.Close()
	if err != nil || err == sql.ErrNoRows {
		log.Fatal(err)
		return nil, 0
	}
	msgMap := make(map[int]ComplaintMsg)
	comMsgStruct := ComplaintMsg{}
	for rows.Next() {
		rows.Scan(&comMsgStruct.Id, &comMsgStruct.Name, &comMsgStruct.Phone, &comMsgStruct.BeComplaint, &comMsgStruct.Reason, &comMsgStruct.Status, &comMsgStruct.DealPerson)
		msgMap[i] = comMsgStruct
		i++
	}

	iterErr := rows.Err() // get any error encountered during iteration
	if iterErr != nil {
		log.Fatal(iterErr)
		return nil, 0
	}
	return msgMap, Num / 10
}

func DoDel(delId int) bool {
	db := getDb()
	defer db.Close()
	query := "delete from complaint where id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer stmt.Close()
	stmt.Exec(delId)
	return true
}

func DoUpdate(id int, status, dealPerson string) bool {
	db := getDb()
	defer db.Close()
	query := "UPDATE complaint SET deal_person=?, status=? WHERE id=?"
	stmt, err := db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	stmt.Exec(dealPerson, status, id)
	return true
}
