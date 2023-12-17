package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 55
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	// another example

	a := 10 // the variable 'a' receive the value 10 in an address

	var pointer *int = &a // the variable 'pointer' receive the address of 'a'
	*pointer = 20         // the variable 'pointer' receive the value 20 in an address of 'a'

	b := &a // the variable 'b' receive the address of 'a'
	*b = 30 // the variable 'b' receive the value 30 in an address of 'a'

	fmt.Println(a)
}
