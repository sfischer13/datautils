// Copyright (c) 2017 Stefan Fischer
// Use of this code is governed by the MIT license, which can be found in the LICENSE file.

package main

import (
	"os"
	"strings"

	"github.com/sfischer13/datautils"

	"github.com/urfave/cli"
)

var (
	date      = "????-??-??"
	goVersion = "go?.?.?"
	version   = "v?.?.?"
)

func main() {
	app := datautils.DefaultApp(version, date, goVersion)
	app.Name = "text"
	app.Usage = "manipulate content from text files"

	app.Commands = []cli.Command{
		{
			Name:   "chars",
			Usage:  "split UTF-8 sequences",
			Action: transformAction(chars),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "words",
			Usage:  "split around white space characters",
			Action: transformAction(words),
			Flags:  []cli.Flag{},
		},
	}

	app.Flags = []cli.Flag{}

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelpAndExit(c, 1)
		return nil
	}

	app.Run(os.Args) // nolint: errcheck, gas
}

func transformAction(function func(string) string) func(*cli.Context) error {
	return func(c *cli.Context) error {
		source := datautils.StdinSource()
		transformer := datautils.TransformPipe(source, function)
		datautils.StdoutSink(transformer)
		return nil
	}
}

func chars(s string) string {
	return strings.Join(strings.Split(s, ""), "\n")
}

func words(s string) string {
	return strings.Join(strings.Fields(s), "\n")
}
