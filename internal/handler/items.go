package handler

import (
	"feedy/internal/models"
	"feedy/internal/store"
	"net/http"
	"strconv"
)

type Items struct {
	Items []*models.Item
}

// GET /items - Get all items
func (h *Handler) HandleGetItems(w http.ResponseWriter, r *http.Request) {
	feedId, err := getFeedIdFromPath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	items, err := store.GetAllItems(h.db, feedId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := Items{
		Items: items,
	}
	h.template.ExecuteTemplate(w, "item_list.html", data)
}

// PATCH /items/:id/read - Toggle item read status
func (h *Handler) HandleToggleItemRead(w http.ResponseWriter, r *http.Request) {
	// Extract item ID
	id, err := getIdFromPath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Toggle read
	if err := store.ToggleRead(h.db, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch updated item and render
	item, err := store.GetItem(h.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.template.ExecuteTemplate(w, "item_row.html", item)
}

// PATCH /items/:id/star - Toggle item star status
func (h *Handler) HandleToggleItemStar(w http.ResponseWriter, r *http.Request) {
	// Extract item ID
	id, err := getIdFromPath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Toggle star
	if err := store.ToggleStar(h.db, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch updated item and render
	item, err := store.GetItem(h.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.template.ExecuteTemplate(w, "item_row.html", item)
}

// Helper function to get the item ID from the request path
func getIdFromPath(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Helper function to get the feed ID from the request path if present
func getFeedIdFromPath(r *http.Request) (int64, error) {
	feedId := r.URL.Query().Get("feed_id")
	if feedId == "" {
		return 0, nil
	}
	id, err := strconv.ParseInt(feedId, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
