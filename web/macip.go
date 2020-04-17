package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/getMac", getMacAddr)
	err := http.ListenAndServe(":6240", nil)
	if err != nil {
		log.Fatal("服务监听异常", err.Error())
	}
}

func getMacAddr(w http.ResponseWriter, r *http.Request) {
	var macAddrs []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		io.WriteString(w, err.Error())
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}
	result, err := json.Marshal(macAddrs)
	io.WriteString(w, string(result))
}
