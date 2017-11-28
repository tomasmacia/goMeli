package domain

import "time"

// Tweet Type Tweet which contains user and text of tweet
type Tweet struct {
	Id   int
	User *User
	Text string
	Date *time.Time
}

// User Type User with its name
type User struct {
	Name     string
	Nick     string
	Email    string
	Password string
}

// NewTweet creates and returns a tweet
func NewTweet(user *User, text string) *Tweet {
	return &Tweet{0, user, text, nil}
}

// NewUser creates and returns a user
func NewUser(username string, nick string, email string, password string) *User {
	return &User{username, nick, email, password}
}
