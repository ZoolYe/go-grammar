package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "使用方式: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, e := net.ResolveTCPAddr("tcp4", service)
	checkError(e)
	conn, e := net.DialTCP("tcp", nil, tcpAddr)
	checkError(e)

	_, e = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(e)

	result, e := ioutil.ReadAll(conn)
	checkError(e)

	fmt.Println(string(result))
	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("未知异常错误", err)
		os.Exit(1)
	}
}
