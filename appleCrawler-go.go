package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq" //why _ ?
	"github.com/line/line-bot-sdk-go/linebot"
)

// postgres:///grimmer

type User struct {
	id   string
	name string
}

type Mac struct {
	SpecsTitle string `json:"specsTitle"`
	SpecsURL   string `json:"specsURL"`
	Price      string `json:"price"`
	// 	ImageURL    string `json:"imageURL"`
	// 	SpecsDetail string `json:"specsDetail"`

}

var bot *linebot.Client

//
var (
	channelID     int64 = 0
	channelSecret       = "0"
	channelMID          = "0"
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
	newMacs, _ := StartCrawer()

	fmt.Println("get new macs callback:", newMacs)

	// for testing
	macs, err := GetAllAppleInfo()
	checkErr(err)

	if reflect.DeepEqual(newMacs, macs) == false {
		//do something
		fmt.Println("Old macs:", macs)

		if len(macs) == 0 {
			fmt.Println("try to insert new macs")
			InsertAppleInfo(newMacs)

		} else {
			fmt.Println("try to update macs")
			UpdateAppleInfo(newMacs)
		}

		// try to broadcast new Macs message
		users, _ := GetAllUserID()
		fmt.Println("get users from db:", users)

		broadcastUpdatedInfo(users, newMacs)

	} else {
		fmt.Println("Same macs")
	}
	// })
}

func broadcastUpdatedInfo(userList []string, macList []Mac) {
	summaryStr := "蘋果特價品更新:" + "\n\n"
	numberOfMacs := len(macList)
	for i, mac := range macList {
		if i == (numberOfMacs - 1) {
			summaryStr = fmt.Sprintf("%s%d. %s. %s  http://www.apple.com%s", summaryStr, i+1, mac.SpecsTitle, mac.Price, mac.SpecsURL)
		} else {
			summaryStr = fmt.Sprintf("%s%d. %s. %s  http://www.apple.com%s", summaryStr, i+1, mac.SpecsTitle, mac.Price, mac.SpecsURL) + "\n\n"
		}
	}
	fmt.Println("new summary macs:", summaryStr)

	fmt.Println("broadcast to:", userList)

	if bot != nil {
		_, err := bot.SendText(userList, summaryStr)

		if err != nil {
			fmt.Println("send fail, ", err)
		}
	}
	fmt.Println("end broadcast to:", userList)

}

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
	var err error
	bot, err = linebot.NewClient(channelID, channelSecret, channelMID)
	checkErr(err)

	go launchCrawer()
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

	// http.HandleFunc("/callback", callbackHandler)
	// port := os.Getenv("PORT")
	// addr := fmt.Sprintf(":%s", port)
	// http.ListenAndServe(addr, nil)

	ticker.Stop()
	fmt.Println("main end")
}
