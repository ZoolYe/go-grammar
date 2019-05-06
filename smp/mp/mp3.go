package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat     int
	progress int
}

func (p *MP3Player) Play(source string) {
	fmt.Println("正在播放MP3", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(10000 * time.Microsecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("播放结束", source)
}
