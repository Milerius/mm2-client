package main

import (
	prompt "github.com/c-bata/go-prompt"
	cli "mm2_client/cli"
)

func main() {
	completer, _ := cli.NewCompleter()
	p := prompt.New(
		cli.Executor,
		completer.Complete,
	)
	p.Run()
}
