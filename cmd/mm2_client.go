package main

import (
	prompt "github.com/c-bata/go-prompt"
	cli "mm2_client/cli"
)

func main() {
	p := prompt.New(
		cli.Executor,
		cli.Completer,
	)
	p.Run()
}
