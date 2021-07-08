package cli

import (
	"fmt"
	"github.com/Delta456/box-cli-maker"
	"github.com/kyokomi/emoji/v2"
	"github.com/manifoldco/promptui"
	"github.com/tyler-smith/go-bip39"
	"mm2_client/config"
	"mm2_client/helpers"
)

func validateRpcPassword(input string) error {
	return nil
}

func generateSeed(cfg *config.MM2Config) {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, err := bip39.NewMnemonic(entropy)

	if err == nil {
		cfg.Passphrase = mnemonic
		cfg.WriteToFile()
		fmt.Printf("MM2.json updated with generated seed: %q\n", mnemonic)
	}
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
		generateSeed(cfg)
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

func CheckMM2Configuration(cfg *config.MM2Config) {
	if cfg.Passphrase == config.NewMM2Config().Passphrase {
		helpers.PrintCheck("Checking if passphrase is configured:", false)
		ConfigurePassPhrase(cfg)
	}

	if cfg.RPCPassword == config.NewMM2Config().RPCPassword {
		helpers.PrintCheck("Checking if rpc password is configured:", false)
		ConfigureRpcPassword(cfg)
	}
}

func CheckMM2DB(cfg *config.MM2Config) {
	prompt := promptui.Select{
		Label: "Do you want to use Desktop DB path ?",
		Items: []string{"Yes", "No"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)

	if result == "Yes" {
		cfg.Dbdir = config.GetDesktopDB()
		cfg.WriteToFile()
		fmt.Printf("Successfully overwrite mm2 db path: %q\n", cfg.Dbdir)
	}
}

func CheckMM2SeedNode(cfg *config.MM2Config) {
	prompt := promptui.Select{
		Label: "Do you want to be a seed node?",
		Items: []string{"Yes", "No"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)

	if result == "Yes" {
		cfg.IMASeed = true
		cfg.WriteToFile()
		fmt.Printf("Successfully overwrite mm2 im_a_seed: %t\n", cfg.IMASeed)
	}
}
