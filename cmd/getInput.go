package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/sorenisanerd/aoc2023/aocclient"
)

type GetInputCommand struct{}

var _ cli.Command = &GetInputCommand{}

func (c GetInputCommand) Run(args []string) int {
	flags := flag.NewFlagSet("getInput", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(c.Help()) }
	year := flags.Int("year", 2023, "Override year")

	if err := flags.Parse(args); err != nil {
		panic(err)
	}

	if flags.NArg() < 1 {
		flags.Usage()
		return 1
	}

	day, err := strconv.Atoi(flags.Arg(0))
	if err != nil {
		panic(err)
	}

	inputReader := aocclient.GetInput(*year, day)

	fp, err := os.Create(fmt.Sprintf("data/%d/day%d.txt", *year, day))
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(fp, inputReader); err != nil {
		panic(err)
	}
	return 0
}

func (c GetInputCommand) Help() string {
	helpText := `
Usage: aoc getInput <DAY>

  This command fetches the input for day DAY of 2023.
`
	return strings.TrimSpace(helpText)
}

func (f GetInputCommand) Synopsis() string {
	return "Fetch input for a given day"
}
