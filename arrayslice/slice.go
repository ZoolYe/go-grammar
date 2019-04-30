package main

import "fmt"

func main() {
	slice2()
}

func slice1() {

	//先定义一个数组
	var myArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于数组创建一个数组切片
	var mySlice = myArray[9:]

	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Println(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _, v := range mySlice {
		fmt.Println(v, " ")
	}

	fmt.Println()
}

func slice2() {

	//创建一个初始元素个数为5的数组切片，元素初始值为0
	//mySlice1 := make([]int, 5)

	//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	//mySlice2 := make([]int, 5, 10)

	//直接创建并初始化包含5个元素的数组切片
	//mySlice3 := []int{1, 2, 3, 4, 5}
}
