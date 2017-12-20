// Copyright (c) 2017 Stefan Fischer
// Use of this code is governed by the MIT license, which can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/sfischer13/datautils"

	"github.com/urfave/cli"
)

var (
	version = "v?.?.?"
	date    = "????-??-??"
)

func main() {
	app := datautils.DefaultApp(version, date)
	app.Name = "rows"
	app.Usage = "select lines from text files"

	app.Commands = []cli.Command{}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rows, r",
			Value: "",
			Usage: "select only lines that are in `LIST`",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		if rows := c.String("rows"); len(rows) != 0 {
			intervals := datautils.String2Intervals(rows)
			function := datautils.Intervals2Func(intervals)
			source := datautils.StdinSource()
			filter := rowFilterPipe(source, function)
			datautils.StdoutSink(filter)
		} else {
			cli.ShowAppHelpAndExit(c, 1)
		}

		return nil
	}

	app.Run(os.Args) // nolint: errcheck, gas
}

func rowFilterPipe(src <-chan string, filter func(int64) bool) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		var i int64
		for s := range src {
			i++
			if filter(i) {
				out <- s
			}
		}
	}()
	return out
}
