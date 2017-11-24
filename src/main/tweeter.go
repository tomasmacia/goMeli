package main

import (
	"github.com/abiosoft/ishell"
	"github.com/goMeli/src/domain"
	"github.com/goMeli/src/service"
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

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			tweet := c.ReadLine()

			service.PublishTweet(domain.NewTweet(user, tweet))

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

			c.Println(tweet.User + ": " + tweet.Text + " (" + tweet.Date.Format("01-02-2006") + ")")

			return
		},
	})

	shell.Run()

}
