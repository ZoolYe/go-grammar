package main

import (
	"errors"
	"fmt"
)

func main() {

}

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("只支持正数的相加")
	}
	return a + b, nil
}

func Add2(a, b int) int {
	return a + b
}

func myFunc(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
	// 按原样传递
	myFunc3(args...)
	//传递片段，实际上任意的int slice都可以传进去
	myFunc3(args[1:]...)
}

func myFunc3(args ...int) {

}
