#!/bin/bash
docker run --rm -v $(pwd):/root ethereum/solc:0.5.7 --abi  /root/$1 -o root/$2
docker run --rm -v $(pwd):/root ethereum/solc:0.5.7 --bin  /root/$1 -o root/$2
#../network_poa/bin/abigen -sol ./f5coin/Business.sol -pkg f5coin -out /src/contract/f5coin.go
