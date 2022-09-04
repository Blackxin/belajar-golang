package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintln(w, "Halo!")
		data := map[string]string{
			"Name":    "Taufik Hidayat",
			"Message": "Hello, world!",
		}

		t, err := template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa kabar!")
}
