package main

import (
	"fmt"
	"runtime"
)

func main() {
	go say("ZoolYe")
	say("Chen")
}

func say(str string) {
	for i := 0; i < 5; i++ {
		//runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		runtime.Gosched()
		fmt.Println(str)
	}
}
