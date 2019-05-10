package main

import "net/http"

const srcDir = "D:/ZOOL/软件"

func main() {
	h := http.FileServer(http.Dir(srcDir))
	http.ListenAndServe(":8081", h)
}
