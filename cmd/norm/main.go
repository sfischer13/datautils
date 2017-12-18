// Copyright (c) 2017 Stefan Fischer
// Use of this code is governed by the MIT license, which can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/sfischer13/datautils"
	"github.com/urfave/cli"

	"golang.org/x/text/unicode/norm"
)

var (
	version = "v?.?.?"
	date    = "????-??-??"
)

func main() {
	app := cli.NewApp()
	app.Name = "norm"
	dateTime := datautils.ParseRFCDate(date).Format("2006-01-02 15:04:05")
	app.Version = fmt.Sprintf("%s (%s)", version, dateTime)
	app.Usage = "normalize text files into Unicode normal forms"
	app.Copyright = "Copyright (c) 2017 Stefan Fischer"
	app.Author = "Stefan Fischer"
	app.Email = "sfischer13@ymail.com"

	app.Commands = []cli.Command{}
	app.HideHelp = true

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "nfc, c",
			Usage: "normalize into NFC",
		},
		cli.BoolFlag{
			Name:  "nfd, d",
			Usage: "normalize into NFD",
		},
		cli.BoolFlag{
			Name:  "nfkc, C",
			Usage: "normalize into NFKC",
		},
		cli.BoolFlag{
			Name:  "nfkd, D",
			Usage: "normalize into NFKD",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		n := datautils.CountTrue([]bool{c.Bool("nfc"), c.Bool("nfd"), c.Bool("nfkc"), c.Bool("nfkd")})
		if n == 0 {
			cli.ShowAppHelpAndExit(c, 1)
		} else if n == 1 {
			if c.Bool("nfc") {
				return datautils.TransformStdin(c, norm.NFC.String)
			} else if c.Bool("nfd") {
				return datautils.TransformStdin(c, norm.NFD.String)
			} else if c.Bool("nfkc") {
				return datautils.TransformStdin(c, norm.NFKC.String)
			} else if c.Bool("nfkd") {
				return datautils.TransformStdin(c, norm.NFKD.String)
			} else {
				return cli.NewExitError("This should not have happened.", 1)
			}
		} else {
			return cli.NewExitError("Only one normal form can be chosen.", 1)
		}

		return nil
	}

	app.Run(os.Args) // nolint: errcheck, gas
}
