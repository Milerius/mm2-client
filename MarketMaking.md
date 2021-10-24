## Linux

```go
// Base coin you want to trade
// "base": "KMD", //< the coin you would like to sell
Base                                  string   `json:"base"`
// Rel coin you want to trade
// "rel": "LTC", //< the coin you would like to get
Rel                                   string   `json:"rel"`
// If you want to trade your whole balance, also usefull if you want to auto adjust amount - optional.
Max                                   bool     `json:"max,omitempty"`
// If you want to trade percentage of your balance - optional and ignored if max field set to true. 
// "balance_percent": "0.5" - range is [0-1]: 0.5 = 50% of the balance
BalancePercent                        string   `json:"balance_percent,omitempty"`
// min_volume in percentage you are willing to accept for your maker order - optional.
// "min_volume": "0.25" - range is [0-1]: 0.25 = 25% of the balance as min_volume
MinVolume                             *string  `json:"min_volume,omitempty"`
//  spread in percentage
// "spread": "1.025" - target_price eg: for KMD/BUSD calculated price 1 (KMD/BUSD price) * 1.025 (+2.5%)
Spread                                string   `json:"spread"`

// confs and nota same as https://developers.komodoplatform.com/basic-docs/atomicdex/atomicdex-api.html#setprice
BaseConfs                             int      `json:"base_confs"`
BaseNota                              bool     `json:"base_nota"`
RelConfs                              int      `json:"rel_confs"`
RelNota                               bool     `json:"rel_nota"`
// The bot will ignore this entry if this value is set to false
Enable                                bool     `json:"enable"`
// will cancel / not create order if price_last_update > price_elapsed_validity (optional fields - default 5min)
PriceElapsedValidity                  *float64 `json:"price_elapsed_validity,omitempty"`
// Will readjust the calculated cex price if a precedent trade exist for the pair or reversed pair - false by default.
// Apply a VWAP logic: https://www.investopedia.com/terms/v/vwap.asp#:~:text=VWAP%20is%20calculating%20the%20sum,periods%20there%20are%20(10).
CheckLastBidirectionalTradeThreshHold *bool    `json:"check_last_bidirectional_trade_thresh_hold,omitempty"`
```

example of configuration that want to swap `KMD/LTC` with `1.5% spread`, `maximum available balance`, `1/4 to be filled`,  
and a `price validity of 30 seconds` (that means that order will be cancelled if the last update from the price 
service is above 30 seconds), with 1 confirmations and without notarization. 
Check trade history with local mm2 DB to never sell < average trading price.

```json
{
    "price_url": "http://price.cipig.net:1313/api/v2/tickers?expire_at=600",
    "cfg": {
        "KMD/LTC": {
            "base": "KMD",
            "rel": "LTC",
            "max": true,
            "min_volume": "0.25",
            "spread": "1.015",
            "base_confs": 1,
            "base_nota": false,
            "rel_confs": 1,
            "rel_nota": false,
            "enable": true,
            "price_elapsed_validity": 30,
            "check_last_bidirectional_trade_thresh_hold": true
        }
    }
}
```

```bash
wget https://github.com/Milerius/mm2-client/releases/download/dev/mm2-tools-client-dev-linux-amd64.tar.gz
tar xvf mm2-tools-client-dev-linux-amd64.tar.gz
./mm2-tools-client
> init
> exit
cd mm2
wget http://195.201.0.6/telegram_bot_notification/mm2-ac768089d-Linux-Release.zip
unzip -o mm2-ac768089d-Linux-Release.zip
# edit mm2_market_maker.json as you wish
cd ..
./mm2-tools-client
> start
> enable_active_coins
> start_simple_market_maker_bot

# When you want to stop (orders can take up to 30 seconds to be cancelled)
> stop_simple_market_maker_bot
```