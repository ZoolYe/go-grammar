package main

import "fmt"

func main() {
	gotoMethod()
}

func ifMethod(x int) int {
	if x < 5 {
		return 0
	} else {
		return 1
	}
}

func switchMethod(i int) {

	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("default")
	}
}

func switchMethod2(num int) {
	switch {
	case 0 <= num && num <= 3:
		fmt.Println("0-3")
	case 4 <= num && num <= 6:
		fmt.Println("4-6")
	}
}

func forMethod1() {
	sum := 0
	for {
		sum++
		if sum > 1000000 {
			break
		}
		fmt.Println("我是循环：", sum)
	}
}

func forMethod2() {
	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func forMethod3() {

	for j := 0; j < 5; j++ {
	JLoop:
		for i := 0; i < 10; i++ {
			if i > 5 {
				//break语句终止的是JLoop标签处的外层循环
				break JLoop
			}
			fmt.Println(i)
		}
	}
}

func gotoMethod() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}
