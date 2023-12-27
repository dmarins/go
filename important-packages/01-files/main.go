package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./import-packages/01-files/file.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Writing data in file"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("File created successfully. Size: %d bytes\n\n", size)

	file, err := os.Open("./import-packages/01-files/file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Printf("File read in parts: %v\n\n", string(buffer[:n]))
	}

	file2, err := os.ReadFile("./import-packages/01-files/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File read in full: %v\n\n", string(file2))
}
