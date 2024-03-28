package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{
					Title: "the GodFather", Director: "Francis bo copolla",
				},
				{
					Title: "Blade Runner", Director: "bradley Scott",
				},
				{
					Title: "The Thing", Director: "John Carpenter",
				},
			},
		}
		tmpl.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	http.HandleFunc("/add-film/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second) // simulate slow server TODO: remove this line on Prod
		log.Print(r.Header.Get("HX-Request"))

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})

	})

	log.Fatal(http.ListenAndServe(":8000", nil))

}
