package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5} // 定义并初始化一个数组

	modify(array) // 传递给一个函数，并试图在函数体内修改这个数组内容

	fmt.Println("In main(), array values:", array)

}

func println(str string) {

	len := len(str)

	for i := 0; i < len; i++ {
		ch := str[i]
		fmt.Println(ch, i)
	}
}

func println2(str string) {
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}

func println3(array []int) {
	for i, v := range array {
		fmt.Println(i, v)
	}
}

func modify(array []int) {
	array[0] = 10 // 试图修改数组的第一个元素
	fmt.Println("In modify(), array values:", array)
}
