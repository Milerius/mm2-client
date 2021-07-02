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

func chooseRpcPassword(cfg *config.MM2Config, hint string) {
	promptChoosePassword := promptui.Prompt{
		Label:    "Please enter a rpc password",
		Validate: validateRpcPassword,
		Default:  hint,
	}

	resultChoose, err := promptChoosePassword.Run()

	thatsFine := true
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		chooseRpcPassword(cfg, "")
		thatsFine = false
	}

	if passErr := helpers.CheckPasswordLever(resultChoose); passErr != nil {
		emoji.Printf("%s\n", passErr.Error())
		chooseRpcPassword(cfg, resultChoose)
		thatsFine = false
	}

	if thatsFine {
		cfg.RPCPassword = resultChoose
		cfg.WriteToFile()
		fmt.Printf("Successfully overwrite mm2 cfg with rpc password: %q\n", resultChoose)
	}
}

func ConfigureRpcPassword(cfg *config.MM2Config) {
	prompt := promptui.Select{
		Label: "RpcPassword Configuration",
		Items: []string{"Choose a rpc password", "Generate a random rpc password"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)

	if result == "Choose a rpc password" {
		chooseRpcPassword(cfg, "")
	} else if result == "Generate a random rpc password" {
		password, _ := helpers.GenerateRandomString(32)
		cfg.RPCPassword = password
		cfg.WriteToFile()
		fmt.Printf("Successfully overwrite mm2 cfg with rpc password: %q\n", password)
	}
}
