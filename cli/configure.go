package cli

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"mm2_client/config"
	"mm2_client/helpers"
)
import "github.com/manifoldco/promptui"

func validateRpcPassword(input string) error {
	return nil
}

func ConfigurePassPhrase(cfg *config.MM2Config) {
	prompt := promptui.Select{
		Label: "Passphrase Configuration",
		Items: []string{"Generate a Seed", "Restore a Seed"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)
}

func chooseRpcPassword(cfg *config.MM2Config) {
	promptChoosePassword := promptui.Prompt{
		Label:    "Please enter a rpc password",
		Validate: validateRpcPassword,
	}
	resultChoose, err := promptChoosePassword.Run()

	if err != nil || resultChoose == "" {
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		} else {
			emoji.Println("Password cannot be empty, please try again :x:")
		}
		chooseRpcPassword(cfg)
	}

	cfg.RPCPassword = resultChoose
	cfg.WriteToFile()
	fmt.Printf("Successfully overwrite mm2 cfg with rpc password: %q\n", resultChoose)
}

func ConfigureRpcPassword(cfg *config.MM2Config) {
	prompt := promptui.Select{
		Label: "RpcPassword Configuration",
		Items: []string{"Choose a rpc password", "Generate a random rpc password"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)

	if result == "Choose a rpc password" {
		chooseRpcPassword(cfg)
	} else if result == "Generate a random rpc password" {
		password, _ := helpers.GenerateRandomString(32)
		cfg.RPCPassword = password
		cfg.WriteToFile()
		fmt.Printf("Successfully overwrite mm2 cfg with rpc password: %q\n", password)
	}
}
