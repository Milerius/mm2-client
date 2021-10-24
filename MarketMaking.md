## Linux

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