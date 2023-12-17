package main

import (
	"errors"
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func getResult(value int) (string, error) {
	if value == 1 {
		return "success", nil
	}

	return "fail,", errors.New("an error occurred")
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	result, error := getResult(0)
	fmt.Println(result, error)

	result, error = getResult(1)
	fmt.Println(result, error)
}
