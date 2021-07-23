package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type CoingeckoHistoryResponse struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Image  struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
	} `json:"image"`
	MarketData struct {
		CurrentPrice struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"current_price"`
		MarketCap struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"market_cap"`
		TotalVolume struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"total_volume"`
	} `json:"market_data"`
	CommunityData struct {
		FacebookLikes            interface{} `json:"facebook_likes"`
		TwitterFollowers         int         `json:"twitter_followers"`
		RedditAveragePosts48H    float64     `json:"reddit_average_posts_48h"`
		RedditAverageComments48H float64     `json:"reddit_average_comments_48h"`
		RedditSubscribers        int         `json:"reddit_subscribers"`
		RedditAccountsActive48H  string      `json:"reddit_accounts_active_48h"`
	} `json:"community_data"`
	DeveloperData struct {
		Forks                        interface{} `json:"forks"`
		Stars                        interface{} `json:"stars"`
		Subscribers                  interface{} `json:"subscribers"`
		TotalIssues                  interface{} `json:"total_issues"`
		ClosedIssues                 interface{} `json:"closed_issues"`
		PullRequestsMerged           interface{} `json:"pull_requests_merged"`
		PullRequestContributors      interface{} `json:"pull_request_contributors"`
		CodeAdditionsDeletions4Weeks struct {
			Additions interface{} `json:"additions"`
			Deletions interface{} `json:"deletions"`
		} `json:"code_additions_deletions_4_weeks"`
		CommitCount4Weeks interface{} `json:"commit_count_4_weeks"`
	} `json:"developer_data"`
	PublicInterestStats struct {
		AlexaRank   int         `json:"alexa_rank"`
		BingMatches interface{} `json:"bing_matches"`
	} `json:"public_interest_stats"`
}

var GeckoPriceAtDateRegistry sync.Map

const (
	layoutCoingecko = "02-01-2006"
	coingeckoUri    = "https://api.coingecko.com/api/v3/coins/"
	geckoPostUri    = "/history?date="
)

func TimestampToGeckoDate(timestamp int64) string {
	var curTime = time.Unix(timestamp, 0)
	return curTime.Format(layoutCoingecko)
}

func HandleGeckoPrice(timestamp int64, geckoId string) {
	uri := coingeckoUri + geckoId + geckoPostUri + TimestampToGeckoDate(timestamp) + "&localization=false"
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusOK {
		//fmt.Printf("OK for %s\n", geckoId)
		defer resp.Body.Close()
		var answer = &CoingeckoHistoryResponse{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			return
		}
		key := geckoId + "-" + TimestampToGeckoDate(timestamp)
		GeckoPriceAtDateRegistry.Store(key, fmt.Sprintf("%.2f", answer.MarketData.CurrentPrice.Usd))
	} else if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Println("Too many request, waiting 1sec and retrying")
		time.Sleep(1 * time.Second)
		HandleGeckoPrice(timestamp, geckoId)
	}
}

func ExistInGeckoRegistry(timestamp int64, geckoId string) bool {
	key := geckoId + "-" + TimestampToGeckoDate(timestamp)
	if _, ok := GeckoPriceAtDateRegistry.Load(key); ok {
		return true
	}
	return false
}

func GetFromRegistry(timestamp int64, geckoId string) string {
	key := geckoId + "-" + TimestampToGeckoDate(timestamp)
	if resp, ok := GeckoPriceAtDateRegistry.Load(key); ok {
		return resp.(string)
	}
	return "0"
}
