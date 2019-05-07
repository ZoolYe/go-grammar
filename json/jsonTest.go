package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     string
	Publisher   string
	IsPublished bool
	Price       float64
}

func main() {
	/*gobook := Book{"Go语言编程", "许式伟",
		"ituring.com.cn", true, 59.29}

	bytes, _ := json.Marshal(gobook)
	fmt.Println(string(bytes))*/

	var book Book

	b := []byte(`{"Title":"Go语言编程","Authors":"许式伟","Publisher":"ituring.com.cn","IsPublished":true,"Price":59.29}`)

	json.Unmarshal(b, &book)

	fmt.Println(book)

}
