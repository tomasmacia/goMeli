package domain

import (
	"time"
)

var nextId int

// TextTweet Type Tweet which contains user and text of tweet
type TextTweet struct {
	Id   int
	User *User
	Text string
	Date *time.Time
}

// ImageTweet Tweets that contain an image
type ImageTweet struct {
	TextTweet
	Image string
}

// QuoteTweet Tweets that contain another Tweet
type QuoteTweet struct {
	TextTweet
	QuotedTweet *TextTweet
}

// Tweeter interface prints tweet
type Tweeter interface {
	PrintableTweet() string
	String() string
}

func giveNextID() int {
	returnID := nextId
	nextId++
	return returnID
}

// NewTweet  NewTweet creates and returns a tweet
func NewTweet(user *User, text string) *TextTweet {
	return NewTextTweet(user, text)
}

// NewTextTweet NewTextTweet creates and returns a text tweet
func NewTextTweet(user *User, text string) *TextTweet {
	return &TextTweet{giveNextID(), user, text, nil}
}

// NewImageTweet NewImageTweet creates and returns a new image tweet
func NewImageTweet(user *User, text string, image string) *ImageTweet {
	return &ImageTweet{TextTweet{giveNextID(), user, text, nil}, image}
}

// NewQuoteTweet NewQuoteTweet creates and returns a new  tweet
func NewQuoteTweet(user *User, text string, tweet *TextTweet) *QuoteTweet {
	return &QuoteTweet{TextTweet{giveNextID(), user, text, nil}, tweet}
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

// PrintableTweet Prints text from tweet
func (qt *QuoteTweet) PrintableTweet() string {
	user := ("@" + qt.User.Nick)
	text := qt.Text
	quotedTweet := qt.QuotedTweet.PrintableTweet()
	return user + ": " + text + " " + "\"" + quotedTweet + "\""
}

// PrintableTweet Prints text from TextTweet
func (tt *TextTweet) String() string {
	return tt.PrintableTweet()
}

// PrintableTweet Prints text from tweet
func (it *ImageTweet) String() string {
	return it.PrintableTweet()
}

// PrintableTweet Prints text from tweet
func (qt *QuoteTweet) String() string {
	return qt.PrintableTweet()
}
