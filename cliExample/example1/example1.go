package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Create a new app
	app := cli.NewApp()
	app.Name = "Your personal greeter"

	// add flags with three arguments
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "stranger",
			Usage: "your wonderful name",
		},
		cli.IntFlag{
			Name:  "age",
			Value: 0,
			Usage: "your graceful age",
		},
	}

	// This function parses and brings data in cli.Context struct
	app.Action = func(c *cli.Context) error {
		// c.String, c.Int for value of given flag
		log.Printf("Hello %s (%d years), Welcome to the command line world", c.String("name"), c.Int("age"))
		return nil
	}

	// Pass os.Args to cli app to parse content
	app.Run(os.Args)
}
