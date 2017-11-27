package service

import (
	"fmt"

	"github.com/curso/goMeli/src/domain"

	"time"
)

var tweets []*domain.Tweet
var nextId int

// InitializeService clears the tweets history
func InitializeService() {
	tweets = tweets[:0]
	nextId = 0
}

// GetTweets returns last tweet
func GetTweets() []*domain.Tweet {
	return tweets
}

// PublishTweet publish tweet
func PublishTweet(tweetToPublish *domain.Tweet) (int, error) {
	var err error
	if tweetToPublish.Text == "" {
		err = fmt.Errorf("text is required")
	}

	nowDate := time.Now()
	tweetToPublish.Date = &nowDate
	tweetToPublish.Id = nextId
	nextId++

	tweets = append(tweets, tweetToPublish)
	return tweetToPublish.Id, err
}

// GetTweetsById dejate de hinchar las bolas compilador
func GetTweetsById(id int) *domain.Tweet {
	for _, v := range tweets {
		if v.Id == id {
			return v
		}
	}
	return nil
}
