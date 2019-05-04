package main

import "fmt"

func main() {

	mySlice := make([]int, 5, 10)

	mySlice = append(mySlice, 1, 2, 3, 4, 5, 6)

	mySlice2 := []int{1, 2, 3, 4, 5}

	//加上省略号相,当于把mySlice2包含的所有元素打散后传入
	// 上述调用等同于 mySlice = append(mySlice, 8, 9, 10)
	mySlice = append(mySlice, mySlice2...)

	fmt.Println("len(mySlice)", len(mySlice))
	fmt.Println("cap(mySlice)", cap(mySlice))

	for _, v := range mySlice {
		fmt.Print(v)
	}

}
