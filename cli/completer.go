package cli

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

var commands = []prompt.Suggest{
	{Text: "exit", Description: "Quit the application"},
	{Text: "init", Description: "Init MM2 Dependencies, Download/Setup"},
	{Text: "help", Description: "Show the global help"},
	{Text: "start", Description: "Start MM2 into a detached process"},
}

var subCommandsHelp = []prompt.Suggest{
	{Text: "init", Description: "Shows help of the init command"},
	{Text: "exit", Description: "Shows help of the help command"},
}

type Completer struct {
}

func NewCompleter() (*Completer, error) {
	return &Completer{}, nil
}

func (c *Completer) argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}
	first := args[0]
	switch first {
	case "help":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommandsHelp, second, true)
		}
	}
	return []prompt.Suggest{}
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	return c.argumentsCompleter(args)
}
