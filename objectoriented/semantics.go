package main

import "fmt"

func main() {
	referenceSemantics()
}

//值语义
func valueSemantics() {
	var a = [3]int{1, 2, 3}
	var b = a
	b[1]++
	fmt.Println(a, b)
}

//引用语义
func referenceSemantics() {
	var a = [3]int{1, 2, 3}

	//赋值语句是数组内容的引用
	//变量b的类型不是[3]int，而是*[3]int类型
	var b = &a

	b[1]++
	fmt.Println(a, *b)
}

/*
	值类型：int、float、bool和string这些类型都属于值类型
	引用类型：特指slice、map、channel这三种预定义类型
*/
