package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq" //why _ ?
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
	fmt.Println("try insert userID")

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
	//  db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	// db, err := sql.Open("postgres", "user=grimmer dbname=grimmer sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	appleInfo, _ := json.Marshal(macInfos)

	stmt, err := db.Prepare("UPDATE special_product_table SET product_info = $1")
	checkErr(err)

	res, err := stmt.Exec(appleInfo)
	checkErr(err)

	// stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	// checkErr(err)
	// res, err := stmt.Exec("astaxieupdate", 1)
	// checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("num of real changed:", affect)

	return err
}

func InsertAppleInfo(macInfos []Mac) error {
	fmt.Println("try insert apple info")
	//  db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	// db, err := sql.Open("postgres", "user=grimmer dbname=grimmer sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client.query(`INSERT INTO special_product_table(product_info) VALUES($1)`,
	//array type of json object : []Mac, need to converted to appleInfo(byte array type)?
	appleInfo, _ := json.Marshal(macInfos)

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

	fmt.Println("num of real changed:", affect)

	return err
}

// may use queryrow
func GetAllAppleInfo() ([]Mac, error) {
	// db, err := sql.Open("postgres", "user=grimmer dbname=grimmer sslmode=disable")
	// if err != nil {
	// 	fmt.Println("connect fail")
	//
	// 	log.Fatal(err) // or panic(err)
	// }
	// defer db.Close()

	fmt.Println("try to GetAllAppleInfo")

	// age := 21
	// rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	rows, err := db.Query("SELECT product_info FROM special_product_table")

	if err != nil {
		fmt.Println("select fail")

		log.Fatal(err)
		// terminate the program ?
		//dial tcp [::1]:5432: getsockopt: connection refused

	}
	defer rows.Close()

	// fmt.Printf("rows:", "(%v, %T)\n", rows, rows)

	var firstMacInfoGroup []Mac
	// var firstMacInfoGroup2 []Mac
	// var firstMacInfoGroup3 []Mac

	for rows.Next() {

		var s []byte

		// if more than one column, do like this, Scan(&name, &age)
		if err := rows.Scan(&s); err != nil {
			fmt.Println("scan error")
			log.Fatal(err)
		}

		// or like https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.2.md
		// another struct, s3 {s2 []Mac }
		var s2 []Mac

		// fmt.Printf("s2:", "(%v, %T)\n", s2, s2)

		// var jsonBlob = []byte(`[
		// 	{"imageURL": "Platypus", "specsURL": "monotremata",
		// 	"specsTitle":"cc", "specsDetail": "Monotremata", "Price":"abc"}
		// ]`)
		err := json.Unmarshal(s, &s2)

		if err != nil {
			fmt.Println("json error:", err)
		} else {

			if firstMacInfoGroup == nil {
				// fmt.Println("before s22", len(s), cap(s))
				// fmt.Println("before first", len(firstMacInfoGroup), cap(firstMacInfoGroup))

				firstMacInfoGroup = s2
				// fmt.Println("assign group:")
				// fmt.Println("afterfirst2", len(firstMacInfoGroup), cap(firstMacInfoGroup))
			}
			// } else if firstMacInfoGroup2 != nil {
			// 	firstMacInfoGroup2 = s2
			// } else if firstMacInfoGroup3 != nil {
			// 	firstMacInfoGroup3 = s2
			// }

			// test update DB
			// if firstMacInfoGroup != nil && firstMacInfoGroup2 != nil {
			// 	fmt.Println("try to compare")
			// 	fmt.Println(reflect.DeepEqual(firstMacInfoGroup, firstMacInfoGroup2))
			//
			// 	firstMacInfoGroup[0].ImageURL = "abc"
			//
			// 	fmt.Println("try to compare again")
			// 	fmt.Println(reflect.DeepEqual(firstMacInfoGroup, firstMacInfoGroup2))
			// }
		}
	}

	// fmt.Println("final")

	return firstMacInfoGroup, err
}
