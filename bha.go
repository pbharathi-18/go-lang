package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func main() {

    templates := template.Must(template.ParseGlob("template/*"))
	http.Handle("/static/",
	http.StripPrefix("/static/",
	   http.FileServer(http.Dir("static"))))
	   http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "homepage.html", nil); err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
 }
 fmt.Println("Listening");
 fmt.Println(http.ListenAndServe(":8080", nil));
}
})