package service

import (
	"fmt"

	"github.com/curso/goMeli/src/domain"

	"time"
)

type TweetManager struct {
	tweets          map[*domain.User][]*domain.Tweet
	nextId          int
	registeredUsers []*domain.User
	loggedUsers     []*domain.User
}

/*
var tweets map[*domain.User][]*domain.Tweet
var nextId int
var registeredUsers []*domain.User
var loggedUsers []*domain.User
*/

// InitializeService clears the tweets history
func (tm *TweetManager) InitializeService() {
	tm.tweets = make(map[*domain.User][]*domain.Tweet)
	tm.registeredUsers = make([]*domain.User, 0)
	tm.loggedUsers = make([]*domain.User, 0)
	tm.nextId = 0
}

// GetTweets returns last tweet
func (tm *TweetManager) GetTweets() []*domain.Tweet {
	var list []*domain.Tweet
	for _, value := range tm.tweets {
		for _, v := range value {
			list = append(list, v)
		}
	}
	return list
}

// PublishTweet publish tweet
func (tm *TweetManager) PublishTweet(tweetToPublish *domain.Tweet) (int, error) {
	var err error
	if tweetToPublish.Text == "" {
		err = fmt.Errorf("text is required")
		return 0, err
	}
	if !tm.isLogged(tweetToPublish.User) {
		err = fmt.Errorf("user not logged in")
		return 0, err
	}

	nowDate := time.Now()
	tweetToPublish.Date = &nowDate
	tweetToPublish.Id = tm.nextId
	tm.nextId++

	value, ok := tm.tweets[tweetToPublish.User]
	if ok {
		tm.tweets[tweetToPublish.User] = append(value, tweetToPublish)
	} else {
		tm.tweets[tweetToPublish.User] = make([]*domain.Tweet, 0)
		tm.tweets[tweetToPublish.User] = append(tm.tweets[tweetToPublish.User], tweetToPublish)
	}

	return tweetToPublish.Id, err
}

// GetTweetsById gets tweets by id
func (tm *TweetManager) GetTweetsById(id int) []*domain.Tweet {
	var list []*domain.Tweet
	for _, value := range tm.tweets {
		for _, v := range value {
			if v.Id == id {
				list = append(list, v)
			}
		}
	}
	return list
}

// GetTweetsByUser gets tweets by user
func (tm *TweetManager) GetTweetsByUser(user *domain.User) []*domain.Tweet {
	return tm.tweets[user]
}

// Register Register one user in service
func (tm *TweetManager) Register(user *domain.User) {
	tm.registeredUsers = append(tm.registeredUsers, user)
}

// Log Log one user in service
func (tm *TweetManager) Log(user *domain.User) error {
	if tm.isRegistered(user) {
		tm.loggedUsers = append(tm.loggedUsers, user)
		return nil
	} else {
		return fmt.Errorf("unregistered user attempted to login")
	}
}

// Logout Logout one user in service
func (tm *TweetManager) Logout(user *domain.User) error {
	if tm.isLogged(user) {
		tm.loggedUsers = deleteFromUserList(tm.loggedUsers, user)
		return nil
	}
	return fmt.Errorf("User is not logged in")
}

func deleteFromUserList(list []*domain.User, user *domain.User) []*domain.User {
	newList := make([]*domain.User, 0)
	for _, v := range list {
		if user != v {
			newList = append(newList, v)
		}
	}
	return newList
}

// GetUsers Get a list of registered users
func (tm *TweetManager) GetUsers() []*domain.User {
	return tm.registeredUsers
}

// GetLoggedUsers Get a list of registered users
func (tm *TweetManager) GetLoggedUsers() []*domain.User {
	return tm.loggedUsers
}

func (tm *TweetManager) isRegistered(user *domain.User) bool {
	for _, v := range tm.registeredUsers {
		if *(user) == *(v) {
			return true
			break
		}
	}
	return false
}

func (tm *TweetManager) isLogged(user *domain.User) bool {
	for _, v := range tm.loggedUsers {
		if *(user) == *(v) {
			return true
			break
		}
	}
	return false
}
