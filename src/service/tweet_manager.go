package service

import (
	"fmt"

	"github.com/curso/goMeli/src/domain"

	"time"
)

// TweetManager Service that manage and publish your tweets
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

	//var tweetCopy *domain.Tweet
	tweetCopy := new(domain.Tweet)
	*tweetCopy = *tweetToPublish

	nowDate := time.Now()
	tweetCopy.Date = &nowDate
	tweetCopy.Id = tm.nextId
	tm.nextId++

	value, ok := tm.tweets[tweetCopy.User]
	if ok {
		tm.tweets[tweetCopy.User] = append(value, tweetCopy)
	} else {
		tm.tweets[tweetCopy.User] = make([]*domain.Tweet, 0)
		tm.tweets[tweetCopy.User] = append(tm.tweets[tweetCopy.User], tweetCopy)
	}

	return tweetCopy.Id, err
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
	userCopy := new(domain.User)
	*userCopy = *user
	tm.registeredUsers = append(tm.registeredUsers, userCopy)
}

// Log Log one user in service
func (tm *TweetManager) Log(user *domain.User) error {
	if tm.isRegistered(user) {
		userCopy := new(domain.User)
		*userCopy = *user
		tm.loggedUsers = append(tm.loggedUsers, userCopy)
		return nil
	}
	return fmt.Errorf("unregistered user attempted to login")
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
		if user.Name != v.Name {
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
		if user.Name == v.Name {
			return true
		}
	}
	return false
}

func (tm *TweetManager) isLogged(user *domain.User) bool {
	for _, v := range tm.loggedUsers {
		if user.Name == v.Name {
			return true
		}
	}
	return false
}

// DeleteTweet Deletes one tweet from every user timeline
func (tm *TweetManager) DeleteTweet(tweetToDelete *domain.Tweet) error {
	// Delete tweet from EVERY timeline, not just from user. TO BE CHECKED
	count := 0
	for user, tweetList := range tm.tweets {
		for _, tweet := range tweetList {
			if tweet.Id == tweetToDelete.Id {
				tm.tweets[user] = deleteFromTweetList(tweetList, tweetToDelete)
				count++

			}
		}
	}
	if count == 0 {
		return fmt.Errorf("Tweet does not exist")
	}
	return nil
}

func deleteFromTweetList(tweetList []*domain.Tweet, tweet *domain.Tweet) []*domain.Tweet {
	newList := make([]*domain.Tweet, 0)
	for _, v := range tweetList {
		if tweet.Id != v.Id {
			newList = append(newList, v)
		}
	}
	return newList
}

// Follow Follow another registered user in the service
func (tm *TweetManager) Follow(userFollowing *domain.User, userFollowed *domain.User) error {
	if !tm.isRegistered(userFollowing) || !tm.isRegistered(userFollowed) {
		return fmt.Errorf("Users must be registered")
	}
	if !tm.isLogged(userFollowing) {
		return fmt.Errorf("Unlogged users cant follow")
	}
	userFollowing.Follow(userFollowed)
	return nil
}

// GetFollowers Get a List of users followed by user
func (tm *TweetManager) GetFollowed(user *domain.User) []*domain.User {
	return user.Following
}

// Unfollow Unfollow one user you followed before in the service
func (tm *TweetManager) Unfollow(userFollowing *domain.User, userFollowed *domain.User) {
	userFollowing.Unfollow(userFollowed)
}

// EditTweet Edits a tweet with the given text
func (tm *TweetManager) EditTweet(tweetToEdit *domain.Tweet, newText string) error {
	count := 0
	for _, tweetList := range tm.tweets {
		for _, tweet := range tweetList {
			if tweet.Id == tweetToEdit.Id {
				tweet.Text = newText
				count++
			}
		}
	}
	if count == 0 {
		return fmt.Errorf("Tweet does not exist")
	}
	return nil
}
