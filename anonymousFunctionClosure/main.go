//匿名函数和闭包
package main

import "fmt"

func main() {

	f := func(x, y int) int {
		return x + y
	}

	fmt.Println(f(2, 3))

}
