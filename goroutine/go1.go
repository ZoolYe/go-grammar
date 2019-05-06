package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		go Add(i, i)
	}

}

func Add(a, b int) {
	fmt.Println(a, b, a+b)
}
