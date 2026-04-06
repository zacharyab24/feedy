// Package fetch provides functions for fetching RSS feeds.
package fetch

import "github.com/mmcdole/gofeed"

// FetchFeed fetches a feed from the given URL and returns it along with any error.
func FetchFeed(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return feed, nil
}
