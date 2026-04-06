package main

import (
	"database/sql"
	"feedy/internal/db"
	"feedy/internal/handler"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Hard coded config constants. Should probably be loaded from a config file.
const port = "8081"
const dbPath = "data/Feedy.db"

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) {
	t.templates.ExecuteTemplate(w, name, data)
}

// Parse and return a new Templates instance.
// Scans the templates directory and partials sub directory for templates, grabs all HTML files
// and parses them into a single template.Template instance.
func newTemplates() *Templates {
	var t *template.Template
	t = template.Must(template.ParseGlob("templates/*.html"))
	t = template.Must(t.ParseGlob("templates/partials/*.html"))
	return &Templates{templates: t}
}

// Open the database and return a pointer to it.
func startDb(path string) *sql.DB {
	db, err := db.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := startDb(dbPath)
	defer db.Close()
	log.Println("Database opened")

	template := newTemplates()
	log.Println("Templates loaded")

	handler := handler.NewHandler(db, template.templates)
	log.Println("Server started on :" + port)
	if err := http.ListenAndServe(":"+port, handler.Route()); err != nil {
		log.Fatal(err)
	}
}
