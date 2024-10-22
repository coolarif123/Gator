package main

import(
	"context"
	"net/http"
	"io"
	"html"
	"encoding/xml"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	var rssFeed *RSSFeed
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return rssFeed, err
	}
	req.Header.Add("User-Agent", "gator")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return rssFeed, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return rssFeed, err
	}

	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return rssFeed, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for _, r := range rssFeed.Channel.Item {
		r.Title = html.UnescapeString(r.Title)
		r.Description = html.UnescapeString(r.Description)
	}

	return rssFeed, nil
}