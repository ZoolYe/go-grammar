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
			fmt.Println("Can't receive")
		}

		fmt.Println("Received back from client: " + reply)
		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
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
