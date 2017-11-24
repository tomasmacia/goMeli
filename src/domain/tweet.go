package domain

import "time"

// Tweet Type Tweet which contains user and text of tweet
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

// NewTweet creates and returns a tweet
func NewTweet(user string, text string) *Tweet {
	return &Tweet{user, text, nil}
}
