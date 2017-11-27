package service

import (
	"fmt"

	"github.com/curso/goMeli/src/domain"

	"time"
)

var tweets []*domain.Tweet

// InitializeService clears the tweets history
func InitializeService() {
	tweets = tweets[:0]
}

// GetTweets returns last tweet
func GetTweets() []*domain.Tweet {
	return tweets
}

// PublishTweet publish tweet
func PublishTweet(tweetToPublish *domain.Tweet) (err error) {
	if tweetToPublish.Text == "" {
		err = fmt.Errorf("text is required")
	}
	nowDate := time.Now()
	tweetToPublish.Date = &nowDate
	tweets = append(tweets, tweetToPublish)
	return
}
