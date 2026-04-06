package handler

import (
	"feedy/internal/fetch"
	"feedy/internal/models"
	"feedy/internal/store"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

// POST /feeds
func (h *Handler) HandleAddFeeds(w http.ResponseWriter, r *http.Request) {
	// Get form values
	url := r.FormValue("url")
	feed, err := fetch.FetchFeed(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	title := feed.Title
	siteLink := feed.Link

	// Create feed
	id, err := store.CreateFeed(h.db, title, url, siteLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.bulkInsertItems(feed, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.updateTemplates(w)
}

// DELETE /feeds/:id
func (h *Handler) HandleDeleteFeeds(w http.ResponseWriter, r *http.Request) {
	// Get feed id from path
	id, err := getIdFromPath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete feed
	err = store.DeleteFeed(h.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.updateTemplates(w)
}

// POST /feeds/refresh - Refresh all feeds
func (h *Handler) HandleRefreshFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := store.GetAllFeeds(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, modelFeed := range feeds {
		feed, err := fetch.FetchFeed(modelFeed.URL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = h.bulkInsertItems(feed, modelFeed.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	items, err := store.GetAllItems(h.db, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := Items{Items: items}
	h.template.ExecuteTemplate(w, "item_list", data)
}

// helper to bulk insert items into the database
func (h *Handler) bulkInsertItems(feed *gofeed.Feed, id int64) error {
	items := make([]models.Item, len(feed.Items))
	for i := range feed.Items {

		var publishedAt time.Time
		if feed.Items[i].PublishedParsed != nil {
			publishedAt = *feed.Items[i].PublishedParsed
		}

		var GUID string
		if feed.Items[i].GUID != "" {
			GUID = feed.Items[i].GUID
		} else {
			GUID = feed.Items[i].Link
		}

		items[i] = models.Item{
			FeedID:      id,
			GUID:        GUID,
			Title:       feed.Items[i].Title,
			Link:        feed.Items[i].Link,
			Content:     feed.Items[i].Content,
			PublishedAt: publishedAt,
			IsRead:      false,
			IsStarred:   false,
			CreatedAt:   time.Now(),
		}
	}

	err := store.BulkCreateItems(h.db, items)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) updateTemplates(w http.ResponseWriter) {
	// Get updated feed and items and update templates
	feedData, err := store.GetAllFeeds(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	itemData, err := store.GetAllItems(h.db, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := HomePageData{
		Feeds: feedData,
		Items: itemData,
	}
	h.template.ExecuteTemplate(w, "feed_list", data)
	h.template.ExecuteTemplate(w, "item_list_oob", data)
}
