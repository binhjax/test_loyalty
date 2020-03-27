#!/bin/bash
# ../network_poa/bin/abigen -abi ./build/loyalty/Business.abi -pkg contracts -out ../src/contracts/Business.go
#../network_poa/bin/abigen -abi ./build/f5coin/Owned.abi -pkg owned -out ../src/contract/owned.go
#../network_poa/bin/abigen -abi ./build/f5coin/Stash.abi -pkg stash -out ../src/contract/stash.go
# ../network_poa/bin/abigen --sol  loyalty/Business.sol  -pkg loyalty  -out ../app/test/loyalty/Business.go
../network_poa/bin/abigen --abi  $1  -pkg contracts  -out $2
