package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := path.Join("views", "index.html")
		tmpl, err := template.ParseFiles(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]any{
			"title": "Learning golang web",
			"name":  "Spider-Man",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
