package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/grimmer0125/stringutil"
)

type Vertex struct {
	xx int
	yy int
}

// import "net/http"

var i = 0

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Foo struct {
	Bar string
}

func testList() {

	// var macs []*Mac
	// printSlice(macs)
	//
	// v := Mac{title: "1"}
	// macs = append(macs, &v)
	// printSlice(macs)
	//
	// newMac := macs[0]
	// fmt.Println("new mac:", *newMac)
	// (*newMac).title = "2"
	//
	// fmt.Println("new mac2:", *newMac)
	//
	// fmt.Println("old mac:", v)
}

func TestFunctions() {

	fmt.Println("teststruct", Vertex{1, 2})

	// const placeOfInterest = `⌘`
	//
	// fmt.Printf("plain string: ")
	// fmt.Printf("%s", placeOfInterest)
	// fmt.Println("change")
	// fmt.Println(placeOfInterest)
	//
	// const placeOfInterest2 = "⌘"
	//
	// fmt.Printf("plain string2: ")
	// fmt.Printf("%s", placeOfInterest2)
	// fmt.Println("change")
	// fmt.Println(placeOfInterest2)
	// fmt.Println("end")

	var j = 10
	i++
	j++
	j++

	fmt.Printf(stringutil.Reverse("!oG ,olleH"))

	fmt.Printf("hello world %s", "grimmer0125")
	fmt.Println("Hello, 世界")

	var aa int = 7

	// aa := testGet()
	aa++

	// resp, err := http.Get("http://example.com/")
	// // fmt.Println("body:", resp.Body)
	// if err != nil {
	// 	fmt.Println("got error")
	// 	// log.Fatal(err)
	// } else {
	// 	defer resp.Body.Close()
	//
	// 	//way1, show ok
	// 	// _, err2 := io.Copy(os.Stdout, resp.Body) //https://gist.github.com/ijt/950790
	// 	//
	// 	// if err2 != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	//
	// 	//way2, 官網的
	// 	//兩進位, 所要要嘛再用string()轉一次, 要嘛用上面的os.Stdout會自動轉成string印出來
	// 	// body, _ := ioutil.ReadAll(resp.Body) //https://golang.org/pkg/net/http/,
	// 	// fmt.Print(body)
	// }

	//way 3 for json
	fmt.Println("START GITHUB")
	r, err := http.Get("https://api.github.com/users/github")
	if err != nil {
		fmt.Println("got error")
		return
	}
	defer r.Body.Close()
	// _, err2 := io.Copy(os.Stdout, r.Body) //https://gist.github.com/ijt/950790
	//
	// if err2 != nil {
	// 	fmt.Println("got error")
	// 	return
	// }

	// foo1 := new(Foo) // or &Foo{}
	// json.NewDecoder(r.Body).Decode(foo1)
	// fmt.Println("get:", foo1.Bar)
	// fmt.Printf(err)

	body, _ := ioutil.ReadAll(r.Body)
	//string()應該是多餘的, 可以直接丟到unmarshal裡面,
	bodyStr := string(body)
	fmt.Printf(bodyStr)

	var f interface{}
	json.Unmarshal([]byte(bodyStr), &f)
	fmt.Println(f)

	testJson()
}

func testJson() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err3 := json.Unmarshal(b, &f)

	if err3 != nil {
		fmt.Println("err3", err3)
	}

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	fmt.Println("final")

}

func testGet() int {
	return 7
	// resp, err := http.Get("http://example.com/")
}
