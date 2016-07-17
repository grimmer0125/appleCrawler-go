package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq" //why _ ?
)

// for debugging, may influence cpu and have side effect
// http://webcache.googleusercontent.com/search?q=cache:2iyEKmfEot0J:colobu.com/2016/04/01/how-to-get-goroutine-id/+&cd=2&hl=zh-TW&ct=clnk&gl=tw
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

// func (s *Selection) Each(f func(int, *Selection)) *Selection {
// 	for i, n := range s.Nodes {
// 		f(i, newSingleSelection(n, s.document))
// 	}
// 	return s
// }

type Mac struct {
	title string
	url   string
	price string
}

func printSlice(s []*Mac) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println("get request")

	// fmt.Println("in handler, its id:", GoID())

}

func launchCrawer() {
	StartCrawer(func(macs []Mac) {
		// fmt.Println("get callback:", macs)
	})
}

// postgres:///grimmer

type Serverslice struct {
	ImageURL    string `json:"imageURL"`
	SpecsURL    string `json:"specsURL"`
	SpecsTitle  string `json:"specsTitle"`
	SpecsDetail string `json:"specsDetail"`
	Price       string `json:"price"`
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func getAllMacInfo() {
	db, err := sql.Open("postgres", "user=grimmer dbname=grimmer sslmode=disable")
	if err != nil {
		log.Fatal(err) // or panic(err)
		fmt.Println("connect fail")

	}

	// age := 21
	// rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	rows, err := db.Query("SELECT product_info FROM special_product_table")

	if err != nil {
		log.Fatal(err)
		fmt.Println("no get rows")

	} else {
		fmt.Println("get rows")
	}

	for rows.Next() {

		fmt.Println("print row")
		var s []byte
		if err := rows.Scan(&s); err != nil {
			log.Fatal(err)
			fmt.Println("scan error")
		}

		// or like https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.2.md
		// another struct, s3 {s2 []Serverslice }
		var s2 []Serverslice

		// var jsonBlob = []byte(`[
		// 	{"imageURL": "Platypus", "specsURL": "monotremata",
		// 	"specsTitle":"cc", "specsDetail": "Monotremata", "Price":"abc"}
		// ]`)
		err := json.Unmarshal(s, &s2)

		if err != nil {
			fmt.Println("json error:", err)
		}

		// fmt.Println("scan ok,s:", s)
		fmt.Println("scan ok,s2:", s2)
	}

	fmt.Println("final")

	//need close

	db.Close()
}

func main() {
	// fmt.Println("in main, its id:", GoID())

	mac1 := Mac{"1", "2", "pp"}
	fmt.Println("mac1:", mac1)

	getAllMacInfo()

	fmt.Println("afte sql")

	launchCrawer()

	ticker := time.NewTicker(time.Second * 60 * 12)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			launchCrawer()
		}
	}()

	fmt.Println("main server start")

	// https://golang.org/doc/articles/wiki/
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("already start server")

	ticker.Stop()

	fmt.Println("main end")
}
