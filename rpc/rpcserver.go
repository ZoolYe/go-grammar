package main

import (
	"errors"
	"fmt"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func main() {

	var arith Arith
	var args = Args{1, 3}
	var reply int

	arith.Multiply(&args, &reply)
	fmt.Println(reply)

	var args2 = Args{9, 3}
	var quotient Quotient
	arith.Divide(&args2, &quotient)
	fmt.Println(quotient.Rem, quotient.Quo)

}

func (arith *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (arith *Arith) Divide(args *Args, quotient *Quotient) error {
	if args.B == 0 {
		return errors.New("不能被0整除")
	}
	quotient.Quo = args.A / args.B
	quotient.Rem = args.B % args.B
	return nil
}
