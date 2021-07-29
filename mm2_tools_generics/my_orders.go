package mm2_tools_generics

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/helpers"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"mm2_client/services"
	"os"
	"runtime"
	"sync"
)

func renderTableTakerOrders(takerOrders map[string]mm2_data_structure.TakerOrderContent) {
	var data [][]string

	for _, cur := range takerOrders {
		var out = []string{cur.Request.Base, cur.Request.BaseAmount, "", cur.Request.RelAmount, cur.Request.Rel, cur.Request.Uuid}
		data = append(data, out)
	}

	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Base", "Base Amount", "", "Rel Amount", "Rel", "UUID"})
		table.SetFooter([]string{"", "", "", "", "", "TakerOrders"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
		fmt.Println()
	} else {
		//goland:noinspection GoPrintFunctions
		fmt.Println("No taker orders\n")
	}
}

func renderTableMakerOrders(withFees bool, makerOrders map[string]mm2_data_structure.MakerOrderContent) {
	var data = make([][]string, len(makerOrders))

	i := 0
	var wg sync.WaitGroup
	for _, cur := range makerOrders {
		wg.Add(1)
		go func(wg *sync.WaitGroup, idx int) {
			defer wg.Done()
			data[idx] = []string{cur.Base, cur.MinBaseVol, cur.AvailableAmount, "", helpers.BigFloatMultiply(cur.AvailableAmount, cur.Price, 8), cur.Rel, cur.Price, cur.Uuid}
			if withFees {
				resp, err := TradePreimage(cur.Base, cur.Rel, cur.Price, "setprice", cur.AvailableAmount)
				if resp != nil {
					if resp.Result != nil {
						fees := ""
						for _, curFee := range resp.Result.TotalFees {
							if helpers.AsFloat(curFee.RequiredBalance) > 0 {
								val, _, _ := services.RetrieveUSDValIfSupported(curFee.Coin)
								if val != "0" {
									val = helpers.BigFloatMultiply(curFee.Amount, val, 2)
								}
								fees += curFee.Amount + " " + curFee.Coin + " (" + val + " $)"
							}
						}
						data[idx] = append(data[idx], fees)
					} else {
						fmt.Println(resp.Error)
					}
				} else {
					fmt.Println(err)
				}
			}
		}(&wg, i)
		i += 1
	}
	wg.Wait()

	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		header := []string{"Base", "Base MinVol", "Base Amount", "", "Rel Amount", "Rel", "Price", "UUID"}
		if withFees {
			header = append(header, "Fees")
		}
		table.SetHeader(header)
		footer := []string{"", "", "", "", "", "", "", "MakerOrders"}
		if withFees {
			footer = append(footer, "")
		}
		table.SetFooter(footer)
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
	} else {
		fmt.Println("No maker orders")
	}
}

func toTableMyOrders(withFees bool, answer *mm2_data_structure.MyOrdersAnswer) {
	renderTableTakerOrders(answer.Result.TakerOrders)
	renderTableMakerOrders(withFees, answer.Result.MakerOrders)
}

func MyOrders() (*mm2_data_structure.MyOrdersAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.MyOrders()
	} else {
		return mm2_http_request.MyOrders()
	}
}

func MyOrdersCLI(withFees bool) {
	if resp, err := MyOrders(); resp != nil {
		toTableMyOrders(withFees, resp)
	} else {
		fmt.Println(err)
	}
}
