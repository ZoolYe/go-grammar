package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player
	switch mtype {
	case "mp3":
		p = &MP3Player{}
	case "wav":
		p = &WAVPlayer{}
	default:
		fmt.Println("未知的音乐类型", mtype)
		return
	}
	p.Play(source)
}
