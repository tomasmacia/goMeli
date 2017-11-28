package service

import (
	"fmt"

	"github.com/curso/goMeli/src/domain"

	"time"
)

var tweets map[*domain.User][]*domain.Tweet
var nextId int

// InitializeService clears the tweets history
func InitializeService() {
	tweets = make(map[*domain.User][]*domain.Tweet)
	nextId = 0
}

// GetTweets returns last tweet
func GetTweets() []*domain.Tweet {
	var list []*domain.Tweet
	for _, value := range tweets {
		for _, v := range value {
			list = append(list, v)
		}
	}
	return list
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

	value, ok := tweets[tweetToPublish.User]
	if ok {
		tweets[tweetToPublish.User] = append(value, tweetToPublish)
	} else {
		tweets[tweetToPublish.User] = make([]*domain.Tweet, 0)
		tweets[tweetToPublish.User] = append(tweets[tweetToPublish.User], tweetToPublish)
	}

	return tweetToPublish.Id, err
}

// GetTweetsById gets tweets by id
func GetTweetsById(id int) []*domain.Tweet {
	var list []*domain.Tweet
	for _, value := range tweets {
		for _, v := range value {
			if v.Id == id {
				list = append(list, v)
			}
		}
	}
	return list
}

// GetTweetsByUser gets tweets by user
func GetTweetsByUser(user *domain.User) []*domain.Tweet {
	return tweets[user]
}
