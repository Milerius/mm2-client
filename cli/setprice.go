package cli

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"mm2_client/helpers"
	"mm2_client/http"
	"strconv"
)

func retrieveCancelPrevious(base string, rel string) bool {
	prompt := promptui.Select{
		Label: "Do you want to cancel previous orders for the selected pair (" + base + "/" + rel + ") ?",
		Items: []string{"Yes", "No"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)

	if result == "Yes" {
		return true
	} else {
		return false
	}
}

func promptMinVol(base string, vol string) string {
	validate := func(input string) error {
		if !helpers.IsNumeric(input) {
			return errors.New("input should be a valid number")
		}
		if helpers.AsFloat(input) > helpers.AsFloat(vol) {
			return errors.New("min_volume must be less or equal than " + vol)
		}
		return nil
	}

	promptAskMinVol := promptui.Prompt{
		Label:    "Please enter your desired min_vol (couldn't be more than " + vol + ")",
		Validate: validate,
	}
	resultMinVol, err := promptAskMinVol.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		promptMinVol(base, vol)
	}
	return resultMinVol
}

func askMinVolume(base string, volume *string, max *bool) *string {
	curVol := "0"
	if volume != nil {
		curVol = *volume
	}
	if max != nil && *max {
		curVol = http.MyBalance(base).Balance
	}
	fmt.Println("The minimum amount of base coin available for the order; it must be less or equal than volume param;")
	curVol = promptMinVol(base, curVol)
	return &curVol
}

func retrieveMinVolume(base string, volume *string, max *bool) *string {
	prompt := promptui.Select{
		Label: "Do you want to use a min_volume for (" + base + ") ?",
		Items: []string{"Yes", "No"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)

	if result == "Yes" {
		return askMinVolume(base, volume, max)
	} else {
		return nil
	}
}

func askRetrieveConfs(ticker string, header string) *int {
	validate := func(input string) error {
		if !helpers.IsInteger(input) {
			return errors.New("input should be a valid integer > 0")
		}
		return nil
	}

	promptAskConfs := promptui.Prompt{
		Label:    "Please enter your desired " + header + " confs for " + ticker,
		Validate: validate,
	}
	resultConfs, err := promptAskConfs.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		askRetrieveConfs(ticker, header)
	}
	val, _ := strconv.Atoi(resultConfs)
	return &val
}

func retrieveConfs(ticker string, header string) *int {
	prompt := promptui.Select{
		Label: "Do you want to use a " + header + " confs for " + ticker,
		Items: []string{"Yes", "No"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)

	if result == "Yes" {
		return askRetrieveConfs(ticker, header)
	} else {
		return nil
	}
}

func retrieveNota(ticker string, header string) *bool {
	prompt := promptui.Select{
		Label: "Do you want to use a " + header + " nota for " + ticker,
		Items: []string{"Yes", "No"},
	}
	_, result, _ := prompt.Run()

	fmt.Printf("You choose %q\n", result)

	if result == "Yes" {
		return helpers.BoolAddr(true)
	} else {
		return nil
	}
}

func interractiveSetPrice(base string, rel string, price string, volume *string, max *bool) {
	cancelPrevious := retrieveCancelPrevious(base, rel)
	minVolume := retrieveMinVolume(base, volume, max)
	baseConfs := retrieveConfs(base, "base")
	relConfs := retrieveConfs(rel, "rel")
	baseNota := retrieveNota(base, "base")
	relNota := retrieveNota(rel, "rel")
	if resp := http.SetPrice(base, rel, price, volume, max, cancelPrevious, minVolume, baseConfs, baseNota, relConfs, relNota); resp != nil {
		resp.ToTable()
	}
}

func SetPrice(base string, rel string, price string, volumeOrMax string) {
	var max *bool = nil
	var volume *string = nil
	if volumeOrMax == "max" {
		max = helpers.BoolAddr(true)
	} else {
		volume = &volumeOrMax
	}
	interractiveSetPrice(base, rel, price, volume, max)
}
