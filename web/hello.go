package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
	http.ResponseWriter类型的对象用于包装处理HTTP服务端的响应信息
	*http.Request表示的是此次HTTP请求的一个数据结构体
*/
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello , World!")
	id := r.URL.RawQuery
	fmt.Println(id)
}

func main() {
	//该方法用于分发请求，即针对某一路径请求将其映射到指定的业务逻辑处理方法中
	http.HandleFunc("/hello", helloHandler)
	/*
		该方法用于在示例中监听 8080 端口，
		接受并调用内部程序来处理连接到此端口的请求。如果端口监听失败，
		会调用log.Fatal()方法输出异常出错信息
	*/
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("服务监听异常", err.Error())
	}
}
