// Package models defines the data models used in the Feedy application.
package models

import "time"

type Feed struct {
	ID        int64
	Title     string
	URL       string
	SiteLink  string
	CreatedAt time.Time
}

type Item struct {
	ID          int64
	FeedID      int64
	GUID        string
	Title       string
	Link        string
	Content     string
	PublishedAt time.Time
	IsRead      bool
	IsStarred   bool
	CreatedAt   time.Time
}
