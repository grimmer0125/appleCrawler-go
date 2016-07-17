package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

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

// var pg_url string

//
// var (
// 	channelID     int64 = 0
// 	channelSecret       = "0"
// 	channelMID          = "0"
// )

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

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	received, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, result := range received.Results {
		content := result.Content()
		if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeText {
			text, err := content.TextContent()
			_, err = bot.SendText([]string{content.From}, "OK "+text.Text)
			if err != nil {
				log.Println(err)
			}
		}
	}
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
	if err != nil {
		fmt.Println("can not open db")
		log.Fatal(err)
	}
	defer db.Close()

	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		db, err = sql.Open("postgres", (db_url + " sslmode=disable"))
	} else {
		db, err = sql.Open("postgres", db_url)
	}

	strID := os.Getenv("channelID")
	numID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		fmt.Println("Wrong environment setting about ChannelID")
	}
	// var err error
	bot, err = linebot.NewClient(numID, os.Getenv("channelSecret"), os.Getenv("channelMID"))
	checkErr(err)

	go launchCrawer()
	ticker := time.NewTicker(time.Second * 60 * 12)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			launchCrawer()
		}
	}()

	fmt.Println("start web server")

	port := os.Getenv("PORT")
	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "5000"
	}
	addr := fmt.Sprintf(":%s", port)

	// https://golang.org/doc/articles/wiki/
	http.HandleFunc("/callback", callbackHandler)
	http.ListenAndServe(addr, nil)
	fmt.Println("already start server")

	// http.HandleFunc("/callback", callbackHandler)
	// port := os.Getenv("PORT")
	// addr := fmt.Sprintf(":%s", port)
	// http.ListenAndServe(addr, nil)

	// ticker.Stop()
	fmt.Println("main end")
}
