package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name     string
	Duration int
}

type Courses []Course

// func ToUpper(value string) string {
// 	return strings.ToUpper(value)
// }

func main() {
	err := template.
		Must(template.
			New("template.html").
			// Funcs(template.FuncMap{"ToUpper": ToUpper}).
			Funcs(template.FuncMap{"Capslock": strings.ToUpper}).
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
