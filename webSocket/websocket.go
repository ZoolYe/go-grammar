package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func echo(ws *websocket.Conn) {

	var err error

	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("无法接受消息")
		}

		fmt.Println("收到客户端消息: " + reply)

		msg := "收到:  " + reply
		fmt.Println("发送消息客户端: " + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("发送失败")
			break
		}
	}

}

func main() {

	http.Handle("/", websocket.Handler(echo))

	if err := http.ListenAndServe(":6244", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
