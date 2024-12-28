package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./learn/template.html")

	type dates struct {
		Title string
		Items []string
		Name  string
	}

	var date []dates = []dates{
		{
			Title: "Go Template Example",
			Items: []string{"Item1,Item2,Item3"},
			Name:  "Arshia",
		},
	}

	err := tmpl.Execute(w, date)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
