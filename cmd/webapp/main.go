package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", ShowHomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("internal/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = temp.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
