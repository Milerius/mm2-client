#!/usr/bin/env python3

import json
import requests

r = requests.get('https://prices.komodo.earth/api/v1/tickers').json()

no_prices = [i for i in r if r[i]["price_provider"] == "unknown"]

r = requests.get('https://api.coinpaprika.com/v1/tickers').json()

returned_paprika_ids = set([i["id"] for i in r])

r = requests.get('https://raw.githubusercontent.com/KomodoPlatform/coins/master/utils/coins_config.json').json()

coins_repo_ids = set([r[i]["coinpaprika_id"] for i in r if i in no_prices])

missing = coins_repo_ids - returned_paprika_ids
missing = list(missing)
missing.sort()

for i in missing:
    print(i)
