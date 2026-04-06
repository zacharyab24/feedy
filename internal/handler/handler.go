// Provides the handler for HTTP requests
package handler

import (
	"database/sql"
	"html/template"
	"net/http"
)

type Handler struct {
	db       *sql.DB
	template *template.Template
}

func NewHandler(db *sql.DB, template *template.Template) *Handler {
	return &Handler{
		db:       db,
		template: template,
	}
}

// Route handles HTTP requests by routing them to the appropriate handler:
// GET /{$} Home page (full payout + sidebar + items)
// GET /items Filtered item list (HTMX fragment)
// PATCH /item/{id}/read (Toggle read/unread)
// PATCH /item/{id}/star (Toggle star/unstar)
// POST /feed (Add a new feed)
// DELETE /feed/{id} (Delete a feed)
// POST /feeds/refresh (Refresh all feeds)
// GET /static/ (Serves static files, CSS, etc)
func (h *Handler) Route() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", h.HandleGetHome)
	mux.HandleFunc("GET /items", h.HandleGetItems)
	mux.HandleFunc("PATCH /items/{id}/read", h.HandleToggleItemRead)
	mux.HandleFunc("PATCH /items/{id}/star", h.HandleToggleItemStar)
	mux.HandleFunc("POST /feeds", h.HandleAddFeeds)
	mux.HandleFunc("DELETE /feeds/{id}", h.HandleDeleteFeeds)
	mux.HandleFunc("POST /feeds/refresh", h.HandleRefreshFeeds)
	mux.HandleFunc("GET /static/", http.FileServer(http.Dir("static")).ServeHTTP)
	return mux
}
