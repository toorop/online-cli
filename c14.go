package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func getC14Commands() []cli.Command {
	return []cli.Command{
		// List platforms
		{
			Name:  "platforms",
			Usage: "Return c14 platforms lits",
			Action: func(c *cli.Context) {
				plateforms, err := online.C14GetPlatforms()
				fmt.Println(err, plateforms)
			},
		},
	}
}
