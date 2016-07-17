package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func GetAllUserID() ([]string, error) {

	rows, err := db.Query("SELECT id FROM user_table")
	checkErr(err)
	defer rows.Close()

	var users []string

	for rows.Next() {
		var id string

		err = rows.Scan(&id)
		checkErr(err)

		users = append(users, id)
	}

	return users, err
}

func InsertUserID(userID string) error {
	fmt.Println("try insert userID:", userID)

	stmt, err := db.Prepare("INSERT INTO user_table(id) VALUES($1)")
	checkErr(err)

	res, err := stmt.Exec(userID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("num of real changed:", affect)

	return nil
}

func UpdateAppleInfo(macInfos []Mac) error {

	fmt.Println("try update")

	appleInfo, _ := json.Marshal(macInfos)

	stmt, err := db.Prepare("UPDATE special_product_table SET product_info = $1")
	checkErr(err)

	res, err := stmt.Exec(appleInfo)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("num of real changed:", affect)

	return err
}

func InsertAppleInfo(macInfos []Mac) error {
	fmt.Println("try insert apple info")

	appleInfo, _ := json.Marshal(macInfos)

	stmt, err := db.Prepare("INSERT INTO special_product_table(product_info) VALUES($1)")
	checkErr(err)

	res, err := stmt.Exec(appleInfo)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("num of real changed:", affect)

	return err
}

// may use queryrow
func GetAllAppleInfo() ([]Mac, error) {

	fmt.Println("try to GetAllAppleInfo")

	rows, err := db.Query("SELECT product_info FROM special_product_table")

	if err != nil {
		fmt.Println("select fail")

		log.Fatal(err)
	}
	defer rows.Close()

	var firstMacInfoGroup []Mac

	for rows.Next() {

		var s []byte

		if err := rows.Scan(&s); err != nil {
			fmt.Println("scan error")
			log.Fatal(err)
		}

		// another way https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.2.md
		var s2 []Mac

		err := json.Unmarshal(s, &s2)

		if err != nil {
			fmt.Println("json error:", err)
		} else {

			if firstMacInfoGroup == nil {

				firstMacInfoGroup = s2
			}
		}
	}

	return firstMacInfoGroup, err
}
