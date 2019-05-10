package main

import "fmt"

func main() {

	ci := make(chan int)

	arrInt := []int{1, 2, 3}
	go sum(arrInt, ci)
	fmt.Println(<-ci)

}

func sum(a []int, c chan int) {

	total := 0
	for _, v := range a {
		total += v
	}
	//将total的值发送到channel中
	c <- total
}
