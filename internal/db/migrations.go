package db

import "database/sql"

const createFeeds = `
CREATE TABLE IF NOT EXISTS feeds (
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	title      TEXT    NOT NULL DEFAULT '',
	url        TEXT    NOT NULL UNIQUE,
	site_link  TEXT    NOT NULL DEFAULT '',
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

const createItems = `
CREATE TABLE IF NOT EXISTS items (
	id           INTEGER PRIMARY KEY AUTOINCREMENT,
	feed_id      INTEGER NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
	guid         TEXT    NOT NULL,
	title        TEXT    NOT NULL DEFAULT '',
	link         TEXT    NOT NULL DEFAULT '',
	content      TEXT    NOT NULL DEFAULT '',
	published_at DATETIME,
	is_read      INTEGER NOT NULL DEFAULT 0,
	is_starred   INTEGER NOT NULL DEFAULT 0,
	created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	UNIQUE(feed_id, guid)
);`

const createIndexes = `
CREATE INDEX IF NOT EXISTS idx_items_published  ON items(published_at DESC);
CREATE INDEX IF NOT EXISTS idx_items_feed_id    ON items(feed_id);
CREATE INDEX IF NOT EXISTS idx_items_is_read    ON items(is_read);
CREATE INDEX IF NOT EXISTS idx_items_is_starred ON items(is_starred);`

func migrate(db *sql.DB) error {
	for _, q := range []string{createFeeds, createItems, createIndexes} {
		if _, err := db.Exec(q); err != nil {
			return err
		}
	}
	return nil
}
