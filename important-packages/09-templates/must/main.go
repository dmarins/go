package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Duration int
}

func main() {
	course := Course{"Golang", 10}
	err := template.
		Must(template.
			New("CourseTemplate").
			Parse("Course: {{.Name}} / Duration: {{.Duration}}h")).
		Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
