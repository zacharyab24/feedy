package store

import (
	"database/sql"
	"feedy/internal/models"
	"log"
)

// BulkCreateItems inserts a batch of items into the database.
func BulkCreateItems(db *sql.DB, items []models.Item) error {
	query := `INSERT OR IGNORE INTO items (feed_id, guid, title, link, content, published_at) VALUES (?, ?, ?, ?, ?, ?)`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, item := range items {
		_, err := tx.Exec(query, item.FeedID, item.GUID, item.Title, item.Link, item.Content, item.PublishedAt)
		if err != nil {
			tx.Rollback()
			log.Printf("Failed to insert item: %v\n", err)
			return err
		}
	}
	return tx.Commit()
}

// GetAllItems retrieves all items from the database.
// Optional feedId can be provided to filter the results.
func GetAllItems(db *sql.DB, feedId int64) ([]models.Item, error) {
	query := `SELECT id, feed_id, guid, title, link, content, published_at, is_read, is_starred, created_at FROM items`
	var args []any

	if feedId > 0 {
		query += ` WHERE feed_id = ?`
		args = append(args, feedId)
	}
	query += ` ORDER BY published_at DESC`

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Printf("Failed to list items: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.FeedID, &item.GUID, &item.Title, &item.Link, &item.Content, &item.PublishedAt, &item.IsRead, &item.IsStarred, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	log.Printf("Found %d items for feed %d\n", len(items), feedId)
	return items, nil
}

// ToggleRead flips the is_read flag for an item in the database.
func ToggleRead(db *sql.DB, itemId int64) error {
	query := `UPDATE items SET is_read = NOT is_read WHERE id = ?`
	_, err := db.Exec(query, itemId)
	if err != nil {
		log.Printf("Failed to toggle read for item %d: %v\n", itemId, err)
	}
	return err
}

// ToggleStar flips the is_starred flag for an item in the database.
func ToggleStar(db *sql.DB, itemId int64) error {
	query := `UPDATE items SET is_starred = NOT is_starred WHERE id = ?`
	_, err := db.Exec(query, itemId)
	if err != nil {
		log.Printf("Failed to toggle star for item %d: %v\n", itemId, err)
	}
	return err
}
