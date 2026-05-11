package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuarios struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("home", func(w http.ResponseWriter, r *http.Request) {
		u := usuarios{
			"João",
			"joao@gmail.com",
		}
		templates.ExecuteTemplate(w, "home.html", u)

	})
	fmt.Println("Servidor rodando na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
