package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	TestFile := "C:/Users/ZOOL/Desktop/uploads/Snipaste_2019-05-08_10-24-55.png"
	infile, inerr := os.Open(TestFile)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("%x  %s\n", md5h.Sum([]byte("")), TestFile)

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("%x  %s\n", sha1h.Sum([]byte("")), TestFile)
	} else {
		fmt.Println(inerr)
		os.Exit(1)
	}
}
