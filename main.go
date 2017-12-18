// Copyright (c) 2017 Stefan Fischer
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

// TransformStdin applies a string transformation function to stdin.
func TransformStdin(c *cli.Context, transform func(string) string) error {
	process := func(s string) error {
		fmt.Println(transform(s))
		return nil
	}
	return processStdin(c, process)
}
