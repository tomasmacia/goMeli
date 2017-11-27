package main

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/curso/goMeli/src/domain"
	"github.com/curso/goMeli/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			username := c.ReadLine()
			user := domain.NewUser(username)

			c.Print("Write your tweet: ")

			tweet := c.ReadLine()

			err := service.PublishTweet(domain.NewTweet(*user, tweet))

			if err != nil {
				fmt.Println(err.Error())
			}

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(tweet.User.Name + ": " + tweet.Text + " (" + tweet.Date.Format("01-02-2006") + ")")

			return
		},
	})

	shell.Run()

}
