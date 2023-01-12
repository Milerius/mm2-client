#!/bin/bash

pkill -f mm2_tools_server_bin
sleep 5
./mm2_tools_server_bin -only_price_service=true
