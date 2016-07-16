package main

import "fmt"

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

func main() {
	// StartCrawer()
	// TestFunctions()
	// testList()
	StartCrawer(func(macs []Mac) {
		fmt.Println("get callback:", macs)
	})
	fmt.Println("main end")
}
