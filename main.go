// Copyright (c) 2017–2021 Stefan Fischer
// Use of this code is governed by the MIT license, which can be found in the LICENSE file.

package datautils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func processStdin(c *cli.Context, process func(string) error) error {
	if c.NArg() != 0 {
		return cli.NewExitError("Command does not take any arguments!", 1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if err := process(line); err != nil {
			return err
		}
	}
	err := scanner.Err()
	return err
}

// DefaultApp TODO
func DefaultApp(version, date, goVersion string) *cli.App {
	app := cli.NewApp()
	dateTime := ParseRFCDate(date).Format("2006-01-02")
	app.Version = fmt.Sprintf("%s (%s, %s)", version, dateTime, goVersion)
	app.Copyright = "Copyright (c) 2017–2021 Stefan Fischer"
	app.Author = "Stefan Fischer"
	app.Email = "sfischer13@ymail.com"
	app.HideHelp = true
	return app
}

// TransformStdin applies a string transformation function to stdin.
func TransformStdin(c *cli.Context, transform func(string) string) error {
	process := func(s string) error {
		fmt.Println(transform(s))
		return nil
	}
	return processStdin(c, process)
}
