package main

import (
	"feedy/internal/db"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Hard coded config constants. Should probably be loaded from a config file.
const port = "8081"
const dbPath = "/data/Feedy.db"

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) {
	t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseFiles("templates/index.html")),
	}
}

func startDb(path string) {
	db, err := db.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func main() {
	startDb(dbPath)
	log.Println("Server started on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
