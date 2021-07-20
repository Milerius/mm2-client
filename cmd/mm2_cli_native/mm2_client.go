package main

import (
	prompt "github.com/c-bata/go-prompt"
	"github.com/kpango/glg"
	cli "mm2_client/cli"
	"mm2_client/constants"
	"mm2_client/log"
	"os"
	"path/filepath"
)

func main() {
	_ = os.MkdirAll(filepath.Join(constants.GetAppDataPath(), "logs"), os.ModePerm)
	infolog := log.InitLogger(filepath.Join(constants.GetAppDataPath(), "logs"), glg.WRITER, "mm2.client")
	defer infolog.Close()
	completer, _ := cli.NewCompleter()
	p := prompt.New(
		cli.Executor,
		completer.Complete,
	)
	p.Run()
}
