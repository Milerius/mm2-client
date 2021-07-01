package cli

import (
	"fmt"
	"mm2_client/config"
)
import "github.com/manifoldco/promptui"

func ConfigurePassPhrase(cfg *config.MM2Config) {
	prompt := promptui.Select{
		Label: "Passphrase Configuration",
		Items: []string{"Generate a Seed", "Restore a Seed"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)
}

func ConfigureRpcPassword(cfg *config.MM2Config) {
	prompt := promptui.Select{
		Label: "RpcPassword Configuration",
		Items: []string{"Choose a rpc password", "Generate a random rpc password"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)
}
