package mm2_wasm_request

import (
	"errors"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
)

func MyTxHistory(coin string, defaultNbTx int, defaultPage int,
	withFiatValue bool, isMax bool) (*mm2_data_structure.MyTxHistoryAnswer, error) {
	return nil, errors.New("not implemented yet")
}
