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

// TextTweet Tweet with just text
type TextTweet struct {
	Tweet
}

// ImageTweet Tweets that contain an image
type ImageTweet struct {
	Tweet
	Image string
}

// Printable interface prints tweet
type Printable interface {
	PrintableTweet() string
}

// NewTweet  NewTweet creates and returns a tweet
func NewTweet(user *User, text string) *Tweet {
	return &Tweet{0, user, text, nil}
}

// NewTextTweet NewTextTweet creates and returns a text tweet
func NewTextTweet(user *User, text string) *TextTweet {
	return &TextTweet{Tweet{0, user, text, nil}}
}

// NewImageTweet NewImageTweet creates and returns a new image tweet
func NewImageTweet(user *User, text string, image string) *ImageTweet {
	return &ImageTweet{Tweet{0, user, text, nil}, image}
}

// PrintableTweet Prints text from TextTweet
func (tt *TextTweet) PrintableTweet() string {
	user := ("@" + tt.User.Nick)
	text := tt.Text
	return user + ": " + text
}

// PrintableTweet Prints text from tweet
func (it *ImageTweet) PrintableTweet() string {
	user := ("@" + it.User.Nick)
	text := it.Text
	image := it.Image
	return user + ": " + text + " " + image
}
