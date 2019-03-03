package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hareku/sitemapaths/pkg/urlpath"
	"github.com/urfave/cli"
	"github.com/yterajima/go-sitemap"
	"golang.org/x/xerrors"
)

func main() {
	app := cli.NewApp()

	app.Name = "sitemapaths"
	app.Usage = "This tool lists paths on sitemap.xml"
	app.Version = "0.1.0"
	app.Action = appAction

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "separator",
			Usage: "separator of paths for output",
			Value: "\n",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("sitemapaths error: %+v", err)
	}
}

func appAction(context *cli.Context) error {
	if context.NArg() == 0 {
		return xerrors.New("must pass sitemap.xml URL to the argument")
	}

	sitemapURL := context.Args()[0]

	sitemapResult, err := sitemap.Get(sitemapURL, nil)
	if err != nil {
		return xerrors.Errorf("fetching sitemap.xml error %q: %w", sitemapURL, err)
	}

	var paths []string
	for _, URL := range sitemapResult.URL {
		path, err := urlpath.GetPath(URL.Loc)
		if err == nil {
			paths = append(paths, path)
		}
	}

	fmt.Print(strings.Join(paths[:], context.String("separator")))

	return nil
}
