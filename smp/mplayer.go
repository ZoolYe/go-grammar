package main

import (
	"bufio"
	"fmt"
	"go-grammar/smp/entry"
	"go-grammar/smp/lib"
	"go-grammar/smp/mp"
	"os"
	"strconv"
	"strings"
)

var mm *lib.MusicManager
var id int = 1
var ctrl, signal chan int

func main() {
	fmt.Println(` 
		输入以下命令来控制播放器: 
		lib list 									  	-- 查看音乐库 
		lib add <歌曲名><歌手><歌曲路径><歌曲文件格式>	-- 添加音乐到音乐库 
		lib remove <数组索引> 						  	-- 从歌曲库中移除此歌曲 
		play <name> 								  	-- 播放音乐
		`)

	mm = lib.NewMusicManager()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("输入指令->： ")
		warLine, _, _ := reader.ReadLine()
		line := string(warLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("无法识别指令:", tokens[0])
		}
	}
}

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < mm.Len(); i++ {
			e, _ := mm.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			mm.Add(&entry.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>", len(tokens))
		}
	case "remove":
		if len(tokens) == 3 {
			index, _ := strconv.Atoi(tokens[2])
			mm.Remove(index)
		} else {
			fmt.Println("USAGE: lib remove <index>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	musicEntry := mm.Find(tokens[1])
	if musicEntry == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}

	mp.Play(musicEntry.Source, musicEntry.Type)

}
