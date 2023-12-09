package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/sorenisanerd/aoc2023/cmd"
)

func commands() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"getInput": func() (cli.Command, error) {
			return cmd.GetInputCommand{}, nil
		},

		"submitAnswer": func() (cli.Command, error) {
			return cmd.SubmitAnswerCommand{}, nil
		},

		"run": func() (cli.Command, error) {
			return cmd.RunCommand{}, nil
		},
	}
}

func main() {
	c := cli.NewCLI("aoc", "2023.12.09")
	c.Args = os.Args[1:]
	c.Commands = commands()
	exitStatus, err := c.Run()

	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
