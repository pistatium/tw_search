package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	ExitCodeOK            int = iota
	ExitCodeArgumentError
	ExitCodeTwitterError
)

func main() {
	app := cli.NewApp()
	app.Name = "tw_search"
	app.Usage = "Twitter searh command"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "access_token, AT",
			Usage:  "Twitter Access Token",
			EnvVar: "TS_AT",
		},
		cli.StringFlag{
			Name:   "access_secret, AS",
			Usage:  "Twitter Access Secret",
			EnvVar: "TS_AS",
		},
		cli.StringFlag{
			Name:   "consumer_key, CK",
			Usage:  "Twitter Consumer Key",
			EnvVar: "TS_CK",
		},
		cli.StringFlag{
			Name:   "consumer_secret, CS",
			Usage:  "Twitter Consumer Secret",
			EnvVar: "TS_CS",
		},
	}

	app.Action = searchAction

	app.Run(os.Args)
}

func searchAction(c *cli.Context) error {
	at := c.GlobalString("access_token")
	as := c.GlobalString("access_secret")
	ck := c.GlobalString("consumer_key")
	cs := c.GlobalString("consumer_secret")

	if at == "false" || as == "false" || ck == "false" || cs == "false" {
		return cli.NewExitError("Twitter OAuth keys must be set.", ExitCodeArgumentError)
	}
	if len(c.Args()) == 0 {
		return cli.NewExitError("Search query must not be blank.", ExitCodeArgumentError)
	}

	query := strings.Join(c.Args(), " ")

	client := twClient(at, as, ck, cs)

	result, err := search(query, client)
	if err != nil {
		return cli.NewExitError(err, ExitCodeTwitterError)
	}
	for i, tweet := range result.Statuses {
		fmt.Printf("\n\n%2d-------\n%v", i+1, tweet.Text)
	}
	return nil
}

func search(query string, client *twitter.Client) (*twitter.Search, error) {
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
	})
	return search, err
}

func twClient(at, as, ck, cs string) *twitter.Client {
	config := oauth1.NewConfig(ck, cs)
	token := oauth1.NewToken(at, as)
	httpClient := config.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}
