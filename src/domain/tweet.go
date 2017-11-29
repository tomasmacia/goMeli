package domain

import (
	"time"
)

// Tweet Type Tweet which contains user and text of tweet
type Tweet struct {
	Id   int
	User *User
	Text string
	Date *time.Time
}

// User Type User with its name
type User struct {
	Name      string
	Nick      string
	Email     string
	Password  string
	Following []*User
}

// NewTweet creates and returns a tweet
func NewTweet(user *User, text string) *Tweet {
	return &Tweet{0, user, text, nil}
}

// NewUser creates and returns a user
func NewUser(username string, nick string, email string, password string) *User {
	return &User{username, nick, email, password, make([]*User, 0)}
}

// Follow Follow another user from service
func (u *User) Follow(user *User) {
	u.Following = append(u.Following, user)
}

// Unfollow Unfollow one user from service
func (u *User) Unfollow(user *User) {
	u.Following = deleteFromUserList(u.Following, user)
}

func deleteFromUserList(userList []*User, user *User) []*User {
	newList := make([]*User, 0)
	for _, u := range userList {
		if user != u {
			newList = append(newList, u)
		}
	}
	return newList
}
