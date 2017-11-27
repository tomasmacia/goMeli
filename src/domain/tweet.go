package domain

import "time"

// Tweet Type Tweet which contains user and text of tweet
type Tweet struct {
	User User
	Text string
	Date *time.Time
}

// User Type User with its name
type User struct {
	Name string
}

// NewTweet creates and returns a tweet
func NewTweet(user User, text string) *Tweet {
	return &Tweet{user, text, nil}
}

// NewUser creates and returns a user
func NewUser(user string) *User {
	return &User{user}
}
