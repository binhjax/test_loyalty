#Project Struture
  ## Network setup: In folder network_poa
    1. Init node data
      # cd network_poa
      #./init_local.sh node1
      #./init_local.sh node2
    2. Start nodes
    => Open 2 console window and run separate command
      #./run_local.sh node1
      #./run_local.sh node2
  ## Code: In folder app
    1. Run docker-compose to initialize environment
      #cd app
      #docker-compose up
    2. Run api server
      #cd server
      #go run main.go
    3. To send request to API or docker-compose
    - Load file app/client/Loyalty.postman_collection.json vào postmain
    4. Document
    - swagger serve -F=swagger swagger.yaml
  ## Contract
    1. Generate contract code if need
    #cd contracts
    #./generate_abi.sh  wallet/Business.sol build/Business
    #./generate_golang.sh build/Business/Business.abi ../app/server/contracts/Business.go
    #Update content in 2 file:
    # - build/Business/Business.abi
    # - build/Business/Business.bin
    # into  app/config/contract.yml
    #
