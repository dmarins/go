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

	template := template.New("CourseTemplate")
	tmp, err := template.Parse("Course: {{.Name}} / Duration: {{.Duration}}h")
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
