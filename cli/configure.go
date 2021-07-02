package cli

import (
	"fmt"
	"github.com/Delta456/box-cli-maker"
	"github.com/kyokomi/emoji/v2"
	"github.com/manifoldco/promptui"
	"mm2_client/config"
	"mm2_client/helpers"
)

func validateRpcPassword(input string) error {
	return nil
}

func restoreSeed(cfg *config.MM2Config) {
	promptRestoreSeed := promptui.Prompt{
		Label: "Please enter your seed",
	}
	resultSeed, err := promptRestoreSeed.Run()
	thatsFine := true
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		restoreSeed(cfg)
		thatsFine = false
	}

	if len(resultSeed) == 0 {
		fmt.Println("You're custom seed cannot be empty, please try again")
		restoreSeed(cfg)
		thatsFine = false
	}

	if thatsFine {
		cfg.Passphrase = resultSeed
		cfg.WriteToFile()
		fmt.Printf("Successfully overwrite mm2 cfg with passphrase: %q\n", resultSeed)
	}
}

func ConfigurePassPhrase(cfg *config.MM2Config) {
	prompt := promptui.Select{
		Label: "Passphrase Configuration",
		Items: []string{"Generate a Seed", "Restore a Seed"},
	}
	_, result, _ := prompt.Run()

	if result == "Generate a Seed" {
		fmt.Printf("You choose %q\n", result)
	} else if result == "Restore a Seed" {
		restoreSeed(cfg)
	}
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
		Box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", Type: "Round", Color: "Green"})
		toPrint := emoji.Sprintf("%s\n", passErr.Error())
		Box.Println("Password Policy", toPrint)
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
