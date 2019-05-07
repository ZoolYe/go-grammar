package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("请求异常")
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
