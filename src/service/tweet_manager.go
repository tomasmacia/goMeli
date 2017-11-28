package service

import (
	"fmt"

	"github.com/curso/goMeli/src/domain"

	"time"
)

var tweets map[*domain.User][]*domain.Tweet
var nextId int
var registeredUsers []*domain.User
var loggedUsers []*domain.User

// InitializeService clears the tweets history
func InitializeService() {
	tweets = make(map[*domain.User][]*domain.Tweet)
	registeredUsers = make([]*domain.User, 0)
	loggedUsers = make([]*domain.User, 0)
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
		return 0, err
	}
	if !isLogged(tweetToPublish.User) {
		err = fmt.Errorf("user not logged in")
		return 0, err
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

// Register Register one user in service
func Register(user *domain.User) {
	registeredUsers = append(registeredUsers, user)
}

// Log Log one user in service
func Log(user *domain.User) error {
	if isRegistered(user) {
		loggedUsers = append(loggedUsers, user)
		return nil
	} else {
		return fmt.Errorf("unregistered user attempted to login")
	}
}

// GetUsers Get a list of registered users
func GetUsers() []*domain.User {
	return registeredUsers
}

// GetLoggedUsers Get a list of registered users
func GetLoggedUsers() []*domain.User {
	return loggedUsers
}

func isRegistered(user *domain.User) bool {
	for _, v := range registeredUsers {
		if *(user) == *(v) {
			return true
			break
		}
	}
	return false
}

func isLogged(user *domain.User) bool {
	for _, v := range loggedUsers {
		if *(user) == *(v) {
			return true
			break
		}
	}
	return false
}
