package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name     string
	Duration int
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := template.
			Must(template.
				New("template.html").
				ParseFiles("template.html")).
			Execute(w,
				Courses{
					{"Golang", 10},
					{"Javascript", 20},
					{"Typescript", 4},
				},
			)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
