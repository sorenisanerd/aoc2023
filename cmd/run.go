package cmd

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
)

type RunCommand struct{}

var _ cli.Command = &RunCommand{}

func (c RunCommand) Run(args []string) int {
	flags := flag.NewFlagSet("RunCommand", flag.ContinueOnError)
	input := flags.String("input", "", "input file")
	flags.Usage = func() { fmt.Println(c.Help()) }
	year := flags.Int("year", 2023, "Override year")

	if len(args) < 1 {
		flags.Usage()
		return 1
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	if input == nil {
		*input = fmt.Sprintf("data/%d/day%d.txt", year, day)
	}

	cmd := exec.Command("go", "run", fmt.Sprintf("./y%d/day%02d", *year, day), *input)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output), err)
		return 1
	}
	fmt.Print(string(output))
	return 0
}

func (c RunCommand) Help() string {
	helpText := `
Usage: aoc run [DAY]

  Run either all days or the given day.
`
	return strings.TrimSpace(helpText)
}

func (f RunCommand) Synopsis() string {
	return "Run given day's solution"
}
