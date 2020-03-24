#!/bin/bash
docker run --rm -v $(pwd):/root ethereum/solc:0.5.7 --abi /root/f5coin/Business.sol -o /root/build/f5coin
#../network_poa/bin/abigen -sol ./f5coin/Business.sol -pkg f5coin -out /src/contract/f5coin.go
