package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func getUserCommands() []cli.Command {
	return []cli.Command{
		// Display user information
		{
			Name:  "info",
			Usage: "Return user info",
			Action: func(c *cli.Context) {
				user, err := online.UserGetInfo()
				dieOnError(err)
				fmt.Println(user)
			},
		},
	}
}
