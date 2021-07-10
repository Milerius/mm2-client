package cli

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"mm2_client/helpers"
	"mm2_client/http"
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
		Label:    "Please enter your desired min_vol",
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
	fmt.Println("The minimum amount of base coin available for the order; it must be less or equal than volume param; " +
		"The following values must be greater than or equal to the min_trading_vol of the corresponding coin")
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

func interractiveSetPrice(base string, rel string, price string, volume *string, max *bool) {
	cancelPrevious := retrieveCancelPrevious(base, rel)
	minVolume := retrieveMinVolume(base, volume, max)
	if resp := http.SetPrice(base, rel, price, volume, max, cancelPrevious, minVolume, nil, nil, nil, nil); resp != nil {
		fmt.Println(resp)
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
