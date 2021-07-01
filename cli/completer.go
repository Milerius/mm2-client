package cli

import "github.com/c-bata/go-prompt"

func Completer(t prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "exit", Description: "Quit the application"},
	}
}
