package main

type ID int

func main() {

	var (
		a string = "anything"
		b int    = 1
		c bool   = false
		d ID     = 10
	)

	d = 20

	println(a, b, c, d)

}
