package service

import (
	"github.com/goMeli/src/domain"

	"time"
)

var tweet *domain.Tweet

// GetTweet returns last tweet
func GetTweet() *domain.Tweet {
	return tweet
}

// PublishTweet publish tweet
func PublishTweet(tweetToPublish *domain.Tweet) {
	tweet = tweetToPublish
	nowDate := time.Now()
	tweet.Date = &nowDate
}
