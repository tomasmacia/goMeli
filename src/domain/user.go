package domain

// User Type User with its name
type User struct {
	Name       string
	Nick       string
	Email      string
	Password   string
	Following  []*User
	Favourites []*TextTweet
}

// NewUser creates and returns a user
func NewUser(username string, nick string, email string, password string) *User {
	return &User{username, nick, email, password, make([]*User, 0), make([]*TextTweet, 0)}
}

// Follow Follow another user from service
func (u *User) Follow(user *User) {
	u.Following = append(u.Following, user)
}

// Unfollow Unfollow one user from service
func (u *User) Unfollow(user *User) {
	u.Following = deleteFromUserList(u.Following, user)
}

// IsFollowing returns a bool if user follows the user send as parameter
func (u *User) IsFollowing(user *User) bool {
	for _, eachuser := range u.Following {
		if eachuser == user {
			return true
		}
	}
	return false
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

// AddFavourite adds a tweet to user's favs list
func (u *User) AddFavourite(tweet *TextTweet) {
	u.Favourites = append(u.Favourites, tweet)
}
