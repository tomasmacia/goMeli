package service

import (
	"fmt"

	"github.com/curso/goMeli/src/domain"

	"time"
)

var tweet *domain.Tweet

// GetTweet returns last tweet
func GetTweet() *domain.Tweet {
	return tweet
}

// PublishTweet publish tweet
func PublishTweet(tweetToPublish *domain.Tweet) (err error) {
	if tweetToPublish.Text == "" {
		err = fmt.Errorf("text is required")
	}
	tweet = tweetToPublish
	nowDate := time.Now()
	tweet.Date = &nowDate
	return
}
