// Copyright (c) 2017 Stefan Fischer
// Use of this code is governed by the MIT license, which can be found in the LICENSE file.

package main

import (
	"os"
	"strings"
	"unicode"

	"github.com/sfischer13/datautils"

	"github.com/urfave/cli"
)

var (
	version = "v?.?.?"
	date    = "????-??-??"
)

func main() {
	app := datautils.DefaultApp(version, date)
	app.Name = "trim"
	app.Usage = "remove white space from text files"

	app.Commands = []cli.Command{}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "left, l",
			Usage: "remove white space at the beginning of lines",
		},
		cli.BoolFlag{
			Name:  "right, r",
			Usage: "remove white space at the end of lines",
		},
		cli.BoolFlag{
			Name:  "top, t",
			Usage: "remove white space lines at the beginning of the input",
		},
		cli.BoolFlag{
			Name:  "bottom, b",
			Usage: "remove white space lines at the end of the input",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		n := datautils.CountTrue([]bool{c.Bool("left"), c.Bool("right"), c.Bool("top"), c.Bool("bottom")})
		if n == 0 {
			cli.ShowAppHelpAndExit(c, 1)
		} else if n == 1 {
			if c.Bool("left") {
				return datautils.TransformStdin(c, trimLeft)
			} else if c.Bool("right") {
				return datautils.TransformStdin(c, trimRight)
			} else if c.Bool("top") {
				source := datautils.StdinSource()
				filter := filterHead(source, datautils.IsNotWhite)
				datautils.StdoutSink(filter)
			} else if c.Bool("bottom") {
				source := datautils.StdinSource()
				filter := filterTail(source, datautils.IsNotWhite)
				datautils.StdoutSink(filter)
			} else {
				return cli.NewExitError("This should not have happened.", 1)
			}
		} else {
			return cli.NewExitError("Only one side can be chosen.", 1)
		}

		return nil
	}

	app.Run(os.Args) // nolint: errcheck, gas
}

func filterHead(src <-chan string, filter func(string) bool) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		flag := false
		for s := range src {
			flag = flag || filter(s)
			if flag {
				out <- s
			}
		}
	}()
	return out
}

func filterTail(src <-chan string, filter func(string) bool) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		cache := datautils.NewList()
		for s := range src {
			cache.Add(s)
			if filter(s) {
				for _, e := range cache.Elements {
					out <- e
				}
				cache.Clear()
			}
		}
	}()
	return out
}

func trimLeft(s string) string {
	return strings.TrimLeftFunc(s, unicode.IsSpace)
}

func trimRight(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}
