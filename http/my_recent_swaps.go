package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/helpers"
	"net/http"
	"strconv"
)

type MyRecentSwapsRequest struct {
	Userpass      string `json:"userpass"`
	Method        string `json:"method"`
	Limit         int    `json:"limit"`
	PageNumber    int    `json:"page_number"`
	MyCoin        string `json:"my_coin,omitempty"`
	OtherCoin     string `json:"other_coin,omitempty"`
	FromTimestamp int64  `json:"from_timestamp,omitempty"`
	ToTimestamp   int64  `json:"to_timestamp,omitempty"`
}

type MyRecentSwapsAnswer struct {
	Result struct {
		FromUuid     string `json:"from_uuid"`
		Limit        int    `json:"limit"`
		Skipped      int    `json:"skipped"`
		Total        int    `json:"total"`
		FoundRecords int    `json:"found_records"`
		PageNumber   int    `json:"page_number,omitempty"`
		TotalPages   int    `json:"total_pages"`
		Swaps        []struct {
			ErrorEvents []string `json:"error_events"`
			Events      []struct {
				Event struct {
					Data struct {
						LockDuration              int    `json:"lock_duration,omitempty"`
						MakerAmount               string `json:"maker_amount,omitempty"`
						MakerCoin                 string `json:"maker_coin,omitempty"`
						MakerCoinStartBlock       int    `json:"maker_coin_start_block,omitempty"`
						MakerPaymentConfirmations int    `json:"maker_payment_confirmations,omitempty"`
						MakerPaymentRequiresNota  bool   `json:"maker_payment_requires_nota,omitempty"`
						MakerPaymentLock          int    `json:"maker_payment_lock,omitempty"`
						MyPersistentPub           string `json:"my_persistent_pub,omitempty"`
						Secret                    string `json:"secret,omitempty"`
						StartedAt                 int    `json:"started_at,omitempty"`
						Taker                     string `json:"taker,omitempty"`
						TakerAmount               string `json:"taker_amount,omitempty"`
						TakerCoin                 string `json:"taker_coin,omitempty"`
						TakerCoinStartBlock       int    `json:"taker_coin_start_block,omitempty"`
						TakerPaymentConfirmations int    `json:"taker_payment_confirmations,omitempty"`
						TakerPaymentRequiresNota  bool   `json:"taker_payment_requires_nota,omitempty"`
						Uuid                      string `json:"uuid,omitempty"`
						TakerPaymentLocktime      int    `json:"taker_payment_locktime,omitempty"`
						TakerPubkey               string `json:"taker_pubkey,omitempty"`
						TxHash                    string `json:"tx_hash,omitempty"`
						TxHex                     string `json:"tx_hex,omitempty"`
						Maker                     string `json:"maker,omitempty"`
						MakerPaymentWait          int    `json:"maker_payment_wait,omitempty"`
						TakerPaymentLock          int    `json:"taker_payment_lock,omitempty"`
						MakerPaymentLocktime      int    `json:"maker_payment_locktime,omitempty"`
						MakerPubkey               string `json:"maker_pubkey,omitempty"`
						SecretHash                string `json:"secret_hash,omitempty"`
						Transaction               struct {
							TxHash string `json:"tx_hash"`
							TxHex  string `json:"tx_hex"`
						} `json:"transaction,omitempty"`
					} `json:"data,omitempty"`
					Type string `json:"type"`
				} `json:"event"`
				Timestamp int64 `json:"timestamp"`
			} `json:"events"`
			MyInfo struct {
				MyAmount    string `json:"my_amount"`
				MyCoin      string `json:"my_coin"`
				OtherAmount string `json:"other_amount"`
				OtherCoin   string `json:"other_coin"`
				StartedAt   int    `json:"started_at"`
			} `json:"my_info"`
			MakerCoin     string   `json:"maker_coin"`
			MakerAmount   string   `json:"maker_amount"`
			TakerCoin     string   `json:"taker_coin"`
			TakerAmount   string   `json:"taker_amount"`
			Gui           string   `json:"gui,omitempty"`
			MmVersion     string   `json:"mm_version"`
			SuccessEvents []string `json:"success_events"`
			Type          string   `json:"type"`
			Uuid          string   `json:"uuid"`
			MyOrderUuid   string   `json:"my_order_uuid,omitempty"`
		} `json:"swaps"`
	} `json:"result"`
}

func NewMyRecentSwapsRequest(limit string, pageNumber string, baseCoin string, relCoin string, from string, to string) *MyRecentSwapsRequest {
	genReq := NewGenericRequest("my_recent_swaps")
	limitNb, err := strconv.Atoi(limit)
	if err != nil {
		limitNb = 50
	}
	pageNumberNb, pageNumberErr := strconv.Atoi(pageNumber)
	if pageNumberErr != nil {
		pageNumberNb = 50
	}
	req := &MyRecentSwapsRequest{Userpass: genReq.Userpass, Method: genReq.Method, Limit: limitNb, PageNumber: pageNumberNb}
	if len(baseCoin) > 0 {
		req.MyCoin = baseCoin
	}
	if len(relCoin) > 0 {
		req.OtherCoin = relCoin
	}
	if len(from) > 0 {
		req.FromTimestamp = helpers.SimpleDateToTimestamp(from)
	}
	if len(to) > 0 {
		req.ToTimestamp = helpers.SimpleDateToTimestamp(to)
	}
	return req
}

func (req *MyRecentSwapsRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func ProcessMyRecentSwaps(limit string, pageNumber string, baseCoin string, relCoin string, from string, to string) *MyRecentSwapsAnswer {
	req := NewMyRecentSwapsRequest(limit, pageNumber, baseCoin, relCoin, from, to).ToJson()
	resp, err := http.Post(GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &MyRecentSwapsAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		return answer
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Err: %s\n", bodyBytes)
	}
	return nil
}
