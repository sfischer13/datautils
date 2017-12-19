// Copyright (c) 2017 Stefan Fischer
// Use of this code is governed by the MIT license, which can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strconv"

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
	app.Name = "count"
	app.Usage = "count lines from text files"

	app.Commands = []cli.Command{}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "keys, k",
			Usage: "sort lines by keys",
		},
		cli.BoolFlag{
			Name:  "counts, c",
			Usage: "sort lines by counts",
		},
		cli.BoolFlag{
			Name:  "reverse, r",
			Usage: "show keys/counts in reverse order",
		},
		cli.BoolFlag{
			Name:  "flip, f",
			Usage: "flip columns",
		},
		cli.BoolFlag{
			Name:  "align, a",
			Usage: "align counts with spaces",
		},
		cli.StringFlag{
			Name:  "delimiter, d",
			Value: "\t",
			Usage: "use `DELIM` between columns",
		},
		cli.IntFlag{
			Name:  "threshold, t",
			Value: 1,
			Usage: "only show counts >= `N`",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		if c.Bool("keys") && c.Bool("counts") {
			return cli.NewExitError("Options --keys and --counts are mutually exclusive!", 1)
		}
		if c.Bool("reverse") && !c.Bool("keys") && !c.Bool("counts") {
			return cli.NewExitError("Option --reverse requires either --keys or --counts!", 1)
		}

		source := datautils.StdinSource()
		counter := rowCounterPipe(source, c)
		datautils.StdoutSink(counter)

		return nil
	}

	app.Run(os.Args) // nolint: errcheck, gas
}

func rowCounterPipe(src <-chan string, c *cli.Context) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)

		// count lines
		counter := pipeCount(src)

		// transform counter
		ps := datautils.Pairs{}
		var max int64
		for _, k := range counter.Keys {
			v := counter.Tally[k]
			max = datautils.Max(max, v)
			ps = append(ps, datautils.Pair{Key: k, Value: v})
		}

		// sort according to options
		sortPairs(c, &ps)

		// print result
		numfmt := "%d"
		if c.Bool("align") {
			numfmt = fmt.Sprintf("%%%dd", len(strconv.FormatInt(max, 10)))
		}
		threshold := int64(c.Int("threshold"))
		for _, pair := range ps {
			if pair.Value < threshold {
				continue
			}
			if c.Bool("flip") {
				out <- fmt.Sprintf("%s%s"+numfmt, pair.Key, c.String("delimiter"), pair.Value)
			} else {
				out <- fmt.Sprintf(numfmt+"%s%s", pair.Value, c.String("delimiter"), pair.Key)
			}
		}
	}()
	return out
}

type counter struct {
	Tally map[string]int64
	Keys  []string
}

func pipeCount(src <-chan string) counter {
	c := counter{
		make(map[string]int64),
		[]string{},
	}
	for s := range src {
		_, ok := c.Tally[s]
		if !ok {
			c.Keys = append(c.Keys, s)
		}
		c.Tally[s]++
	}
	return c
}

func sortPairs(c *cli.Context, ps *datautils.Pairs) {
	if c.Bool("reverse") && c.Bool("keys") {
		ps.ReverseKeys()
	}
	if c.Bool("reverse") && c.Bool("counts") {
		ps.SortKeys()
		ps.ReverseValues()
	}
	if !c.Bool("reverse") && c.Bool("keys") {
		ps.SortKeys()
	}
	if !c.Bool("reverse") && c.Bool("counts") {
		ps.SortKeys()
		ps.SortValues()
	}
}
