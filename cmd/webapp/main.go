package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nyae44/url-shortener-app/internal/controllers"
	"github.com/nyae44/url-shortener-app/internal/db"
	"log"
	"net/http"
)

func main() {
	sqlite, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite.Close()

	if err := db.CreateTable(sqlite); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			controllers.ShowIndex(w, r)
		} else {
			controllers.Proxy(sqlite)(w, r)
		}
	})
	http.HandleFunc("/shorten", controllers.Shorten(sqlite))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
