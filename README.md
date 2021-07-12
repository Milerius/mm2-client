# mm2-client
my own mm2 client / tooling

Future tasks:

- [ ] Wallet + encryption seed
- [ ] cancel_all_order
- [x] setprice
- [ ] buy
- [ ] sell  
- [x] my_recent_swaps
- [x] my_orders
- [x] prompt use DB desktop
- [x] cancel (if i want to cancel by UUID)
- [x] update_maker_order (want to track UUID)
- [ ] gecko price service
- [ ] paprika price service
- [x] binance websocket service
- [x] add total in my_balance_all
- [x] add am_i_seed in MM2.json if user wants it  
- [x] get_binance_supported_pairs
- [x] add a way to start mm2 without extra services
- [x] generic price service that use in order (binance, gecko, paprika)
- [x] simple bot cfg

## How to use the trading bot on linux:

```
wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
git clone https://github.com/Milerius/mm2-client
cd mm2-client && go build cmd/mm2_client.go
mkdir -p mm2
cp assets/simple_market_bot.template.json mm2/simple_market_bot.json
# edit the cfg if you want and remove the commentary
./mm2_client

> help
> init # you need to do only the first time
> start
> enable_active_coins
> enable COIN_FROM_CFG COIN_2_FROM_CFG # use that if they are not enable yet - you can use active coins next run
> get_binance_supported_pairs COIN_FROM_CFG # you can see if your coin is supported

# Be sure your balance is funded before starting the bot
> start_simple_market_maker_bot

> my_orders

# later
> stop_simple_market_maker_bot
> stop
> exit

## from another terminal
tail -f  ~/.atomicdex_cli/logs/simple.market.maker.logs
```
