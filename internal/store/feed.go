// Package store provides CRUD operations for the Feedy application database.
package store

import (
	"database/sql"
	"feedy/internal/models"
	"log"
)

// CreateFeed creates a new feed in the database and returns the last insert id.
func CreateFeed(db *sql.DB, title string, url string, siteLink string) (int64, error) {
	query := `INSERT INTO feeds (title, url, site_link) VALUES (?, ?, ?)`
	res, err := db.Exec(query, title, url, siteLink)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	log.Printf("Feed created with id: %d\n", lastInsertId)
	return lastInsertId, nil

}

// GetFeed retrieves a feed by id from the database and stores it in the provided models.Feed struct.
func GetFeed(db *sql.DB, id int64) (*models.Feed, error) {
	var f models.Feed
	query := `SELECT id, title, url, site_link, created_at FROM feeds WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&f.ID, &f.Title, &f.URL, &f.SiteLink, &f.CreatedAt)
	if err != nil {
		log.Printf("Failed to get feed with id: %d\n", id)
	} else {
		log.Printf("Feed retrieved with id: %d\n", id)
	}
	return &f, err
}

// GetAllFeeds retrieves all feeds from the database and returns them as a slice of models.Feed structs.
func GetAllFeeds(db *sql.DB) ([]*models.Feed, error) {
	query := `SELECT id, title, url, site_link, created_at FROM feeds`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feeds []*models.Feed
	for rows.Next() {
		f := &models.Feed{}
		if err := rows.Scan(&f.ID, &f.Title, &f.URL, &f.SiteLink, &f.CreatedAt); err != nil {
			return nil, err
		}
		feeds = append(feeds, f)
	}
	return feeds, nil
}

// DeleteFeed deletes a feed by id from the database.
func DeleteFeed(db *sql.DB, id int64) error {
	query := `DELETE FROM feeds WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete feed with id: %d\n", id)
		return err
	}
	log.Printf("Feed deleted with id: %d\n", id)
	return nil
}
