package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	CopyFile("C:/Users/ZOOL/Desktop/hosts", "C:/Users/ZOOL/Documents/hosts1")
}

func CopyFile(dst, src string) (w int64, err error) {

	srcFile, err := os.Open(src)

	if err != nil {
		fmt.Println("在" + src + "路径没有找到相关的文件")
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}

	defer func() {
		dstFile.Close()
		srcFile.Close()
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	foo()
	return io.Copy(dstFile, srcFile)

}

func foo() {
	panic("404")
}
