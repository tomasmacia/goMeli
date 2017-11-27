package service_test

import (
	"testing"

	"github.com/curso/goMeli/src/domain"
	"github.com/curso/goMeli/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	var user *domain.User

	user = domain.NewUser("grupoEsfera")
	text := "This is my first tweet"

	tweet = domain.NewTweet(*user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweets()[0]

	if publishedTweet.User != *user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			*user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetsTwoDifferentsWithoutTweets(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet
	var user *domain.User

	user = domain.NewUser("grupoEsfera")
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(*user, text)
	secondTweet = domain.NewTweet(*user, secondText)

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, user, text) {
		t.Error("First tweet has incorrect information")
		return
	}

	if !isValidTweet(t, secondPublishedTweet, user, secondText) {
		t.Error("Second tweet has incorrect information")
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, user *domain.User, text string) bool {
	return tweet.User == *user && tweet.Text == text
}
