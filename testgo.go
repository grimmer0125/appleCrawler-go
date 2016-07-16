package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/grimmer0125/stringutil"
)

// import "net/http"

var i = 0

func main() {

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
	resp, err := http.Get("http://example.com/")
	// fmt.Println("body:", resp.Body)
	if err != nil {
		fmt.Println("got error")
		// log.Fatal(err)
	} else {
		defer resp.Body.Close()
		_, err2 := io.Copy(os.Stdout, resp.Body) //https://gist.github.com/ijt/950790

		if err2 != nil {
			log.Fatal(err)
		}

		//兩進位
		// body, _ := ioutil.ReadAll(resp.Body) //https://golang.org/pkg/net/http/,
		// fmt.Print(body)
	}

	// fmt.Printf(err)

}

func testGet() int {
	return 7
	// resp, err := http.Get("http://example.com/")
}
