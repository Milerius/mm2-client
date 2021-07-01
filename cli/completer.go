package cli

import "github.com/c-bata/go-prompt"

func Completer(t prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "exit", Description: "Quit the application"},
		{Text: "init", Description: "Init MM2 Dependencies, Download/Setup"},
		{Text: "help <command>", Description: "Show the help for a specific command"},
		{Text: "help", Description: "Show the global help"},
	}
}
