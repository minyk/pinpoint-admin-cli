package main

import (
	"github.com/minyk/pinpoint-cli/cli"
	"gopkg.in/alecthomas/kingpin.v3-unstable"
)

func main() {
	app := cli.New()
	cli.HandleDefaultSections(app)
	kingpin.MustParse(app.Parse(cli.GetArguments()))
}
