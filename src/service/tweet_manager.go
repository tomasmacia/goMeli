package service

var tweet string

// GetTweet returns last tweet
func GetTweet() string {
	return tweet
}

// PublishTweet publish tweet
func PublishTweet(tweetToPublish string) {
	tweet = tweetToPublish
}
