package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq" //why _ ?
)

func InsertAppleInfo(macInfos []MacInDB) error {
	fmt.Println("try insert")
	//  db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	db, err := sql.Open("postgres", "user=grimmer dbname=grimmer sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// client.query(`INSERT INTO special_product_table(product_info) VALUES($1)`,
	//array type of json object : []MacInDB, need to converted to appleInfo(byte array type)?
	appleInfo, _ := json.Marshal(macInfos)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO special_product_table(product_info) VALUES($1)")
	checkErr(err)

	res, err := stmt.Exec(appleInfo)
	checkErr(err)

	// stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	// checkErr(err)
	// res, err := stmt.Exec("astaxieupdate", 1)
	// checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	return nil
}

// may use queryrow
func GetAllAppleInfo() ([]MacInDB, error) {
	fmt.Println("in another file")
	db, err := sql.Open("postgres", "user=grimmer dbname=grimmer sslmode=disable")
	if err != nil {
		fmt.Println("connect fail")

		log.Fatal(err) // or panic(err)
	}
	defer db.Close()

	fmt.Println("try to query")

	// age := 21
	// rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	rows, err := db.Query("SELECT product_info FROM special_product_table")

	if err != nil {
		fmt.Println("select fail")

		log.Fatal(err)
		// terminate the program ?
		//dial tcp [::1]:5432: getsockopt: connection refused

	} else {
		fmt.Println("get rows")
	}

	// fmt.Printf("rows:", "(%v, %T)\n", rows, rows)

	var fistMacInfoGroup []MacInDB
	for rows.Next() {

		fmt.Println("rows.next")

		var s []byte

		// if more than one column, do like this, Scan(&name, &age)
		if err := rows.Scan(&s); err != nil {
			log.Fatal(err)
			fmt.Println("scan error")
		}

		// or like https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.2.md
		// another struct, s3 {s2 []MacInDB }
		var s2 []MacInDB

		fmt.Printf("s2:", "(%v, %T)\n", s2, s2)

		// var jsonBlob = []byte(`[
		// 	{"imageURL": "Platypus", "specsURL": "monotremata",
		// 	"specsTitle":"cc", "specsDetail": "Monotremata", "Price":"abc"}
		// ]`)
		err := json.Unmarshal(s, &s2)

		if err != nil {
			fmt.Println("json error:", err)
		} else {
			fmt.Println("scan ok,s2:")

			if fistMacInfoGroup == nil {
				fmt.Println("assign group:", fistMacInfoGroup)
				fistMacInfoGroup = s2
			}
		}
	}

	fmt.Println("final")

	return fistMacInfoGroup, err
}
