package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tw_search"
	app.Usage = "Twitter searh command"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "access_token, AT",
			Usage:  "Twitter Access Token",
			EnvVar: "TS_AT",
		},
		cli.BoolFlag{
			Name:   "access_secret, AS",
			Usage:  "Twitter Access Secret",
			EnvVar: "TS_AS",
		},
		cli.BoolFlag{
			Name:   "consumer_key, CK",
			Usage:  "Twitter Consumer Key",
			EnvVar: "TS_CK",
		},
		cli.BoolFlag{
			Name:   "consumer_secret, CS",
			Usage:  "Twitter Consumer Secret",
			EnvVar: "TS_CS",
		},
	}

	app.Action = search

	app.Run(os.Args)
}

func search(c *cli.Context) {

	at := c.GlobalString("access_token")
	as := c.GlobalString("access_secret")
	ck := c.GlobalString("consumer_key")
	cs := c.GlobalString("consumer_secret")

	var msg = ""
	if len(c.Args()) > 0 {
		msg = c.Args().First()
		msg += at
		msg += as
		msg += ck
		msg += cs
	}

	fmt.Printf("%s\n", msg)
}
