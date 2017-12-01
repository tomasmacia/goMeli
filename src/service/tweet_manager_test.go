package service_test

import (
	"fmt"
	"testing"

	"github.com/curso/goMeli/src/domain"
	"github.com/curso/goMeli/src/service"
)

var userTest *domain.User = domain.NewUser("grupoEsfera", "grupoesfera", "ge@hotmail.com", "123456")
var tweetTest *domain.Tweet = domain.NewTweet(userTest, "Quiquiriqui")

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var manager service.TweetManager
	manager.InitializeService()

	var user *domain.User
	user = domain.NewUser("grupoEsfera", "asd", "asd", "asd")
	manager.Register(user)
	manager.Log(user)

	text := "This is my first tweet"

	var tweet *domain.Tweet
	tweet = domain.NewTweet(user, text)

	// Operation
	manager.PublishTweet(tweet)

	// Validation
	fmt.Println("DEBUG")
	fmt.Println(manager.GetTweets())
	fmt.Println("DEBUG")
	publishedTweet := manager.GetTweets()[0]

	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			*user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutTextisNotPublished(t *testing.T) {

	// Initialization
	var manager service.TweetManager
	manager.InitializeService()

	var tweet *domain.Tweet
	var user *domain.User

	user = domain.NewUser("grupoEsfera", "asd", "asd", "asd")
	manager.Register(user)
	manager.Log(user)

	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = manager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Errorf("Expected error is text is required")
	}
}

func TestTweetsTwoDifferentsWithoutTweets(t *testing.T) {

	// Initialization
	var manager service.TweetManager
	manager.InitializeService()

	var tweet, secondTweet *domain.Tweet
	var user *domain.User

	user = domain.NewUser("grupoEsfera", "asd", "asd", "asd")
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	manager.Register(user)
	manager.Log(user)

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	manager.PublishTweet(tweet)
	manager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := manager.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstPublishedTweet.Id, user, text) {
		t.Error("First tweet has incorrect information")
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondPublishedTweet.Id, user, secondText) {
		t.Error("Second tweet has incorrect information")
		return
	}

}

func TestCanRetrieveTweetsById(t *testing.T) {

	// Inicializacion
	var manager service.TweetManager
	manager.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := domain.NewUser("grupoEsfera", "asd", "asd", "asd")
	text := "This is my first tweet"

	manager.Register(user)
	manager.Log(user)

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = manager.PublishTweet(tweet)

	// Validation
	publishedTweet := manager.GetTweetsById(id)[0]

	isValidTweet(t, publishedTweet, 0, user, text)
}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user *domain.User, text string) bool {
	return tweet.User == user && tweet.Text == text && tweet.Id == id
}

func TestTweetsByUser(t *testing.T) {

	// Initialization
	var manager service.TweetManager
	manager.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := domain.NewUser("grupoEsfera", "user", "email", "passwd")
	anotherUser := domain.NewUser("nick", "asd", "asd", "asd")

	manager.Register(user)
	manager.Register(anotherUser)
	manager.Log(user)
	manager.Log(anotherUser)

	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	firstId, _ := manager.PublishTweet(tweet)
	secondId, _ := manager.PublishTweet(secondTweet)
	manager.PublishTweet(thirdTweet)

	// Operation
	tweets := manager.GetTweetsByUser(user)
	fmt.Print(*tweets[0])

	// Validation
	if len(tweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}
	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}
}

func TestRegisterOneUser(t *testing.T) {
	// Initialization
	var manager service.TweetManager
	manager.InitializeService()

	manager.Register(userTest)

	if len(manager.GetUsers()) != 1 {
		t.Error("Registered users number must be 1")
	}
}

func TestLoginOneUser(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeService()

	manager.Register(userTest)
	manager.Log(userTest)

	if len(manager.GetLoggedUsers()) != 1 {
		t.Error("Logged users number must be 1")
	}

	manager.PublishTweet(tweetTest)
	tweets := manager.GetTweetsByUser(userTest)
	if len(tweets) != 1 {
		t.Error("Tweet by logued user wasn't registered correctly")
	}
}

func TestLogFail(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeService()

	err := manager.Log(userTest)
	if err == nil {
		t.Error("User not registered cannot log in")
	}
}

func TestTweetFailed(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeService()

	manager.Register(userTest)

	if len(manager.GetUsers()) != 1 {
		t.Error("Registered users number must be 1")
	}

	if len(manager.GetLoggedUsers()) != 0 {
		t.Error("Logged users number must be 0")
	}

	_, err := manager.PublishTweet(tweetTest)
	if err == nil {
		t.Error("Unlogued user was able to tweet")
	}

	tweets := manager.GetTweetsByUser(userTest)
	if len(tweets) != 0 {
		t.Error("Tweet by unlogged user was apparently registered")
	}
}

func TestLogoutOneUser(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeService()

	manager.Register(userTest)
	manager.Log(userTest)

	if len(manager.GetLoggedUsers()) != 1 {
		t.Error("Logged users number must be 1 before logging out")
		return
	}

	manager.Logout(userTest)

	if len(manager.GetLoggedUsers()) != 0 {
		t.Error("Logged users number must be 0 after logging out")
	}
}

func TestLogoutOneUnloggedUser(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeService()

	manager.Register(userTest)

	if manager.Logout(userTest) == nil {
		t.Error("User is not logged before logging in")
	}
}

func TestRemoveOnePublishedTweet(t *testing.T) {
	// Initialization
	var manager service.TweetManager
	manager.InitializeService()

	manager.Register(userTest)
	manager.Log(userTest)

	// Operation
	manager.PublishTweet(tweetTest)
	manager.DeleteTweet(tweetTest)

	// Validation
	publishedTweets := manager.GetTweets()

	if len(publishedTweets) != 0 {
		t.Errorf("Expected tweets should be 0 after removing but is %d", len(publishedTweets))
	}

}

func TestOneUserCanFollowAnotherOne(t *testing.T) {
	// Initialization
	var manager service.TweetManager
	manager.InitializeService()

	userFollowed := domain.NewUser("bot", "bot", "bot", "bot")

	manager.Register(userTest)
	manager.Register(userFollowed)

	manager.Log(userTest)
	manager.Log(userFollowed)

	manager.Follow(userTest, userFollowed)

	if len(manager.GetFollowed(userTest)) != 1 {
		t.Errorf("Expected followers should be 1 but is %d", len(manager.GetFollowed(userTest)))
	}
}

func TestOneUserCantFollowUnregisteredUsers(t *testing.T) {
	// Initialization
	var manager service.TweetManager
	manager.InitializeService()

	userFollowed := domain.NewUser("bot", "bot", "bot", "bot")

	manager.Register(userTest)

	manager.Log(userTest)

	err := manager.Follow(userTest, userFollowed)

	if err == nil {
		t.Errorf("Unregistered users cant be followed")
	}
}

func TestOneUserUnregisteredCantFollowRegisteredUsers(t *testing.T) {
	// Initialization
	var manager service.TweetManager
	manager.InitializeService()

	userFollowed := domain.NewUser("bot", "bot", "bot", "bot")

	manager.Register(userFollowed)

	err := manager.Follow(userTest, userFollowed)

	if err == nil {
		t.Errorf("Unregistered users cant follow users")
	}
}

func TestOneUnloggedUserCantFollowUsers(t *testing.T) {
	// Initialization
	var manager service.TweetManager
	manager.InitializeService()

	userFollowed := domain.NewUser("bot", "bot", "bot", "bot")

	manager.Register(userTest)

	err := manager.Follow(userTest, userFollowed)

	if err == nil {
		t.Errorf("Unlogged users cant follow")
	}
}

func TestOneUserUnfollowsAFollowedUser(t *testing.T) {
	// Initialization
	var manager service.TweetManager
	manager.InitializeService()
	//
	userTest2 := domain.NewUser("sad", "asd", "asd", "asdasd")
	//

	userFollowed := domain.NewUser("bot", "bot", "bot", "bot")

	manager.Register(userTest2)
	manager.Register(userFollowed)
	manager.Log(userTest2)

	// Operation
	manager.Follow(userTest2, userFollowed)

	if len(manager.GetFollowed(userTest2)) != 1 {
		t.Errorf("Followers number should be 1 but is %d", len(manager.GetFollowed(userTest2)))
	}

	manager.Unfollow(userTest2, userFollowed)

	if len(manager.GetFollowed(userTest2)) != 0 {
		t.Errorf("Expected users followed by %v was 0 but is %d", userTest2.Name, len(manager.GetFollowed(userTest2)))
	}
}
func TestEditTweet(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeService()
	manager.Register(userTest)
	manager.Log(userTest)
	manager.PublishTweet(tweetTest)
	manager.EditTweet(tweetTest, "Changed")
	changedTweet := manager.GetTweets()[0]

	if changedTweet.Text != "Changed" {
		t.Errorf("Tweet's text wasn't changed, should be Changed, is %v", changedTweet.Text)
	}
}

func TestReferences(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeService()
	manager.Register(userTest)
	manager.Log(userTest)
	manager.PublishTweet(tweetTest)
	manager.EditTweet(tweetTest, "Changed")

	if tweetTest.Text == "Changed" {
		t.Errorf("Original tweet from tweet_manager_test changed when it shouldn't have")
	}
}

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet(userTest, "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

/*
func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image", "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
*/
