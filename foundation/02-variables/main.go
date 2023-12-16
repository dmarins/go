package main

func main() {

	var anything1 string
	println(anything1)

	var anything2 string = "anything2"
	println(anything2)

	anything3 := "anything3"
	println(anything3)

	const anything4 string = "anything4"
	println(anything4)

	anything5, anything6 := "anything5", "anything6"
	println(anything5, anything6)

	var (
		anything7 string = "anything7"
		anything8 string = "anything8"
	)
	println(anything7, anything8)

}
