package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	client := &http.Client{}
	req.Header.Set("User-Agent", "gator")
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()
	dat, err := io.ReadAll(res.Body)

	var feed *RSSFeed
	err = xml.Unmarshal(dat, &feed)
	if err != nil {
		return &RSSFeed{}, err
	}
	feed.unescapeString()
	return feed, nil
}

func (feed *RSSFeed) unescapeString() {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
}
