//数据类型
package main

import "fmt"

func main() {

	var a bool = true
	fmt.Println(a)

	/*
		var b *int
		var c []int
		var d map[string] int
		var e chan int
		var f func(string) int
		var g error //error 是接口
	*/

	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fn, mn, ln, cc := getName()
	fmt.Println(fn, mn, ln, cc)

	_, _, aa, _ := getName()
	fmt.Println(aa)

}

func getName() (firstName, middleName, lastName, nickName string) {
	firstName = "May"
	middleName = "M"
	lastName = "Chen"
	return
}


