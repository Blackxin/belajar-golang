package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]any

func main() {
	// tmpl, err := template.ParseGlob("views/*")
	// if err != nil {
	// 	panic(err.Error())
	// }

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := M{"name": "Batman"}
		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		err := tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := M{"name": "Spider-Man"}
		tmpl := template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))

		err := tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
