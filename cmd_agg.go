package main

import(
	"context"
	"time"
	"github.com/coolarif123/Gator/internal/database"
	"database/sql"
	"github.com/google/uuid"
	"errors"
	"fmt"
)

func Agg (s *State, cmd Command) error {
	if (len(cmd.Args) < 3 ) {
		return errors.New("not enough arguements input") 
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[2])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err = ScrapeFeeds(s)
		if (err != nil) {
			return err
		}
	}
	return nil
}

func ScrapeFeeds(s *State) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	newMarkFeedFetched := database.MarkFeedFetchedParams{
		LastFetchedAt:	sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		UpdatedAt:		time.Now(),
		ID:				nextFeed.ID,
	}

	err = s.Db.MarkFeedFetched(context.Background(), newMarkFeedFetched)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, r := range rssFeed.Channel.Item{
		urlExists, err := s.Db.UrlExists(context.Background(), r.Link) 
		if (urlExists == 1){
			continue;
		}

		newPost := database.CreatePostParams {
			ID:         uuid.New(),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Title:      sql.NullString{
				String: r.Title,
				Valid:  true,
			},
			Url:         r.Link,
			Description: sql.NullString{
				String: r.Description,
				Valid:	true,
			},
			PublishedAt: r.PubDate,
			FeedID:      nextFeed.ID,
		}

		err = s.Db.CreatePost(context.Background(), newPost)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", r.Title)
	}
	return nil
}


