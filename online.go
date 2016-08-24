package main

import (
	"log"
	"os"

	"github.com/toorop/gonline"
	"github.com/urfave/cli"
)

var online *gonline.Online

func main() {
	var err error
	// init onlineClient
	apiKey := os.Getenv("ONLINE_API_KEY")
	if apiKey == "" {
		log.Fatalln("ONLINE_API_KEY is missing in env")
	}
	online, err = gonline.New(apiKey)
	dieOnError(err)

	app := cli.NewApp()
	app.Name = "online"
	app.Usage = "CLI tool to interact with online.net API"
	app.Author = "Stéphane Depierrepont"
	app.Email = "toorop@toorop.fr"
	app.Version = "0.0.0"

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Commands = []cli.Command{
		{
			Name:        "user",
			Usage:       "user section",
			Subcommands: getUserCommands(),
		}, {
			Name:        "c14",
			Usage:       "c14 section",
			Subcommands: getC14Commands(),
		},
	}
	app.Run(os.Args)
}
