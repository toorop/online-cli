package main

import (
	"fmt"

	"github.com/toorop/gonline"
	"github.com/urfave/cli"
)

func getC14Commands() []cli.Command {

	return []cli.Command{
		// List platforms
		{
			Name:  "platforms",
			Usage: "Return c14 platforms list",
			Action: func(c *cli.Context) {
				plateforms, err := online.C14GetPlatforms()
				dieOnError(err)
				for _, p := range plateforms {
					fmt.Println(p)
				}
			},
		}, {
			Name:        "platform",
			Usage:       "online platform PLATFORM_ID",
			Description: "Return c14 platform detailst",
			Action: func(c *cli.Context) {
				//plateformID := c.Args().First()

				plateform, err := online.C14GetPlatformDetails(1)
				dieOnError(err)

				fmt.Println(plateform)
			},
		}, {
			Name:        "protocols",
			Usage:       "online protocol",
			Description: "Return available protocols",
			Action: func(c *cli.Context) {
				protocols, err := online.C14GetProtocols()
				dieOnError(err)
				for _, p := range protocols {
					fmt.Println(p)
				}
			},
		}, {
			Name:        "safes",
			Usage:       "online safes",
			Description: "Return user's safes",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "count",
					Value: 0,
					Usage: "Pagination: Number of results",
				},
				cli.IntFlag{
					Name:  "stop",
					Value: 0,
					Usage: "Pagination: Maximum id",
				},
				cli.IntFlag{
					Name:  "start",
					Value: 0,
					Usage: "Pagination: Minimum id",
				},
			},
			Action: func(c *cli.Context) {
				config := gonline.C14GetSafesOptions{
					Count: c.Int("count"),
					Start: c.Int("start"),
					Stop:  c.Int("stop"),
				}
				safes, err := online.C14GetSafes(config)
				dieOnError(err)
				for _, s := range safes {
					fmt.Println(s)
				}
			},
		},
	}
}
