package main

import (
	"fmt"
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
	// maye change to goroutine instead of using callback
	StartCrawer(func(macs []Mac) {
		// fmt.Println("get callback:", macs)
	})
}

// postgres:///grimmer

type MacInDB struct {
	ImageURL    string `json:"imageURL"`
	SpecsURL    string `json:"specsURL"`
	SpecsTitle  string `json:"specsTitle"`
	SpecsDetail string `json:"specsDetail"`
	Price       string `json:"price"`
}

// type Serverslice struct {
// 	ImageURL    string
// 	SpecsURL    string
// 	SpecsTitle  string
// 	SpecsDetail string
// 	Price       string
// }

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("err:", err)
		// panic(err)
	}
}

func main() {

	// for testing
	// MacInfoGroup, err := GetAllAppleInfo()
	// checkErr(err)
	// InsertAppleInfo(MacInfoGroup)
	// fmt.Println("after reading, duplicate from sql")

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
