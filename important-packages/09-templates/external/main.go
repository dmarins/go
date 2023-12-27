package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Duration int
}

type Courses []Course

func main() {
	err := template.
		Must(template.
			New("template.html").
			ParseFiles("template.html")).
		Execute(os.Stdout,
			Courses{
				{"Golang", 10},
				{"Golang", 15},
				{"Javascript", 20},
				{"Typescript", 4},
			},
		)
	if err != nil {
		panic(err)
	}
}
