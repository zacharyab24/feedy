package handler

import (
	"feedy/internal/models"
	"feedy/internal/store"
	"net/http"
)

type HomePageData struct {
	Feeds []*models.Feed
	Items []*models.Item
}

// GET / - Home page
func (h *Handler) HandleGetHome(w http.ResponseWriter, r *http.Request) {
	feeds, err := store.GetAllFeeds(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	items, err := store.GetAllItems(h.db, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := HomePageData{
		Feeds: feeds,
		Items: items,
	}
	h.template.ExecuteTemplate(w, "layout.html", data)
}
