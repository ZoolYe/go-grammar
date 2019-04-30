//接口
package main

import "fmt"

type Bird struct {

}

func main() {
	var fly IFly = new(Bird)
	fly.Fly()
}

func (b *Bird) Fly() {
	fmt.Println("我是实现的方法")
}

type IFly interface {
	Fly()
}