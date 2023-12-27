package main

import (
	"fmt"
)

func main()  {

	fmt.Println("line 1")
	defer fmt.Println("line 2")
	fmt.Println("line 3")
	
}