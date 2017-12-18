// Copyright (c) 2017 Stefan Fischer
// Use of this code is governed by the MIT license, which can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sfischer13/datautils"

	"github.com/urfave/cli"
)

var (
	version = "v?.?.?"
	date    = "????-??-??"
)

func main() {
	app := cli.NewApp()
	app.Name = "text"
	dateTime := datautils.ParseRFCDate(date).Format("2006-01-02 15:04:05")
	app.Version = fmt.Sprintf("%s (%s)", version, dateTime)
	app.Usage = "manipulate content from text files"
	app.Copyright = "Copyright (c) 2017 Stefan Fischer"
	app.Author = "Stefan Fischer"
	app.Email = "sfischer13@ymail.com"

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
	app.HideHelp = true

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
