package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	var b Integer = 2
	var a Integer = 1
	//a.Less(b)
	a.Add(b)
	fmt.Println(a)
}

func (a *Integer) Add(b Integer) {
	*a += b
}
