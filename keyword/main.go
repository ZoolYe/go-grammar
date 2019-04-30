//关键字
package main

import "fmt"

func main() {

	//保留的关键字
	var reserved = "break	default	func	interface	select case	defer	go	map	struct chan	else	goto	package	switch const	fallthrough	if	range	type continue	for	import	return	var"

	fmt.Println(reserved)

	//预定义标识符
	var predefinedStandard = "append	bool	byte	cap	close	complex	complex64	complex128	uint16 copy	false	float32	float64	imag	int	int8	int16	uint32 int32	int64	iota	len	make	new	nil	panic	uint64 print	println	real	recover	string	true	uint	uint8	uintptr"

	fmt.Println(predefinedStandard)
}
