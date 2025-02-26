package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Create a cli application
	app := cli.NewApp()

	// Define name and description of app
	app.Name = "Store Marks"
	app.Description = "Your simple marks Store Cli App"

	// define flags to parse
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:     "save",
			Required: true,
			Usage:    "Should save to database (yes/no)",
		},
	}

	app.Version = "1.0"
	// Define Actions
	app.Action = func(c *cli.Context) error {
		var args []string
		if c.NArg() > 0 {
			// Fetch arguments in a array
			args = c.Args()
			personName := args[0]
			marks := args[1:]
			log.Println("Person: ", personName)
			log.Println("marks: ", marks)
		}

		// Check the flag value
		if !c.Bool("save") {
			log.Println("Skipping saving to the database")
		} else {
			// Add database logic here
			log.Println("Saving to the database", args)
		}
		return nil
	}
	app.Run(os.Args)
}
