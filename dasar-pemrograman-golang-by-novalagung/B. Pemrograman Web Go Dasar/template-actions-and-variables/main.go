package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:    "Bruce Wayne",
			Gender:  "male",
			Hobbies: []string{"Reading books", "Travelling", "Buying things"},
			Info:    Info{"Wayne Enterprises", "Gothan City"},
		}

		tmpl := template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
