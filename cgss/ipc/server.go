package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {

	session := make(chan string, 0)

	go func(c chan string) {
		for {
			request := <-c
			if request == "close" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("请求格式无效", request)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b)
		}
		fmt.Println("会话结束")
	}(session)
	fmt.Println("已成功创建新会话")
	return session
}
