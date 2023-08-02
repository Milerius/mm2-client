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
- [x] gecko price service
- [x] paprika price service
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
cd mm2-client && go build cmd/mm2_cli_native/mm2_client.go
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
tail -f ~/.atomicdex_cli/logs/mm2.client.log
```

### How to use the simple market maker bot with an existing atomicdex instance

```
go build -o mm2_tools_server_bin cmd/mm2_tools_server/mm2_tools_server.go
./mm2_tools_server_bin

# Assuming your userpass for the session is foobar
# Starting the simple market maker bot
curl --location --request POST 'localhost:13579/api/v1/start_simple_market_maker_bot' \
--header 'Content-Type: application/json' \
--data-raw '{
  "desktop_cfg_path": "/Users/milerius/coins/utils/coins_config.json",
  "mm2_coins_cfg_path": "/Users/milerius/Library/Application Support/AtomicDex Desktop/0.5.0/configs/coins.json",
  "market_maker_cfg_path": "/Users/milerius/GolandProjects/mm2-client/mm2/simple_market_bot.json",
  "mm2_userpass": "foobar"
}'

# stopping the bot
curl --location --request POST 'localhost:13579/api/v1/stop_simple_market_maker_bot'
```

### How to use the server on Ios

#### building

```
cd mm2_tools_server
gomobile bind -v --target=ios .
```

#### using in an ios project:
```obj-c
//
//  main.m
//  FooBar
//
//  Created by Sztergbaum Roman on 15/07/2021.
//

#import <UIKit/UIKit.h>
#import "AppDelegate.h"
#import "Mm2_tools_server.h"

int main(int argc, char * argv[]) {
    NSString * appDelegateClassName;
    @autoreleasepool {
        // Setup code that might create autoreleased objects goes here.
        appDelegateClassName = NSStringFromClass([AppDelegate class]);
    }
    dispatch_async(dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_DEFAULT, 0), ^{
        Mm2_tools_serverLaunchServer(@"atomic_dex");

    });

    return UIApplicationMain(argc, argv, nil, appDelegateClassName);
}
```

### using in an android project:

#### building

```
cd mm2_tools_server
gomobile bind -v --target=android .
```

#### Using in an android-studio (kotlin) project:

```kt
import mm2_tools_server.Mm2_tools_server
import kotlin.concurrent.thread

class MainActivity : AppCompatActivity() {

    private lateinit var binding: ActivityMainBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        binding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(binding.root)

        val navView: BottomNavigationView = binding.navView

        val navController = findNavController(R.id.nav_host_fragment_activity_main)
        // Passing each menu ID as a set of Ids because each
        // menu should be considered as top level destinations.
        val appBarConfiguration = AppBarConfiguration(
            setOf(
                R.id.navigation_home, R.id.navigation_dashboard, R.id.navigation_notifications
            )
        )

        thread {
            Mm2_tools_server.launchServer("atomicDex")
        }
        //print("hello world\n")
        setupActionBarWithNavController(navController, appBarConfiguration)
        navView.setupWithNavController(navController)
    }
}
```

#### miscs
```
#you may want for testing purpose to forward localhost port of the server
# for android simulator devices please start the emulator and run then start your app
adb forward tcp:1313 tcp:1313
```
