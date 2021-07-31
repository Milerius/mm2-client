package market_making

import (
	"fmt"
	"github.com/kpango/glg"
	"mm2_client/constants"
	"mm2_client/external_services"
	"mm2_client/mm2_tools_generics"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"time"
)

var gPrecedentRecentSwapsRegistry = make(map[string]mm2_data_structure.SwapContent)
var gQuitRecentSwapsUpdated chan struct{}

func processRecentSwapsUpdate() {
	_ = glg.Info("process recent swaps update")
	recentSwaps, err := mm2_tools_generics.MyRecentSwaps("10", "1", "", "", "", "")
	if err != nil {
		_ = glg.Errorf("error my_recent_swaps: %v", err)
		return
	}

	currentSwapsRegistry := recentSwaps.ToMap()
	for uuid, swap := range currentSwapsRegistry {
		if swapPrevious, ok := gPrecedentRecentSwapsRegistry[uuid]; ok {
			if swapPrevious.GetLastStatus() == "Finished" {
				continue
			}
			if swapPrevious.GetLastStatus() != swap.GetLastStatus() {
				glg.Infof("Swap status changed: %s", swap.Uuid)
				external_services.SendMessage("Swap Status changed:", swap.ToMessage())
			}
		} else {
			glg.Infof("New swap started: %s", swap.Uuid)
			external_services.SendMessage("New Swap started: ", swap.ToMessage())
		}
	}
	gPrecedentRecentSwapsRegistry = currentSwapsRegistry
}

func startRecentSwapsUpdate() {
	recentSwaps, err := mm2_tools_generics.MyRecentSwaps("10", "1", "", "", "", "")
	if err != nil {
		_ = glg.Errorf("error my_recent_swaps: %v", err)
		return
	}
	gPrecedentRecentSwapsRegistry = recentSwaps.ToMap()
	for {
		select {
		case <-gQuitRecentSwapsUpdated:
			glg.Info("My recent swaps updater successfully stopped")
			constants.GMyRecentSwapsUpdaterServiceRunning = false
			close(gQuitRecentSwapsUpdated)
			return
		default:
			processRecentSwapsUpdate()
			time.Sleep(5 * time.Second)
		}
	}
}

func bootstrapRecentSwapsUpdate() {
	if !constants.GMM2Running {
		_ = glg.Warn("Couldn't start my recent swaps updater because mm2 is not running")
		return
	}
	if constants.GMyRecentSwapsUpdaterServiceRunning {
		_ = glg.Warn("My Recent Swaps updated service is already running (or being stopped) - skipping")
		return
	}
	_ = glg.Infof("Starting my recent swaps updater")
	constants.GMyRecentSwapsUpdaterServiceRunning = true
	gQuitRecentSwapsUpdated = make(chan struct{})
	go startRecentSwapsUpdate()

}

func stopRecentSwapsUpdate() {
	if constants.GMyRecentSwapsUpdaterServiceRunning {
		//! Also need to cancel all existing orders (Could use by UUID meanwhile this time)
		_ = glg.Info("Stopping My recent swaps updater - may take up to 5 seconds")
		go func() {
			gQuitRecentSwapsUpdated <- struct{}{}
		}()
	} else {
		fmt.Println("My recent swaps updater is still shutting down or not running")
		_ = glg.Warn("My recent swaps updater is still shutting down or not running")
	}
}
