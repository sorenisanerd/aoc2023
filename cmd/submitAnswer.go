package cmd

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/sorenisanerd/aoc2023/aocclient"
	"github.com/sorenisanerd/aoc2023/utils"
)

type SubmitAnswerCommand struct{}

var _ cli.Command = &SubmitAnswerCommand{}

func (c SubmitAnswerCommand) Run(args []string) int {
	flags := flag.NewFlagSet("submitAnswer", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(c.Help()) }
	year := flags.Int("year", 2023, "Override year")

	if len(args) < 3 {
		flags.Usage()
		return 1
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	if day < 0 || day > 25 {
		panic("Invalid day")
	}

	fmt.Println(aocclient.SubmitAnswer(*year, day, utils.MustAtoi(args[1]), args[2]))

	return 0
}

func (c SubmitAnswerCommand) Help() string {
	helpText := `
Usage: aoc submitAnswer <DAY> <PART> <ANSWER>

  This command submits the answer for part PART of day DAY.
`
	return strings.TrimSpace(helpText)
}

func (f SubmitAnswerCommand) Synopsis() string {
	return "Submit answer for a given part of a given day"
}
