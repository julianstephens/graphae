package main

import (
	"github.com/alecthomas/kong"
	"github.com/julianstephens/graphae/cli"
)

func main() {
	cli := cli.CLI{
		Globals: cli.Globals{
			Version: cli.VersionFlag("0.1.1"),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("graphae"),
		kong.Description("An image processing CLI"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		})
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
