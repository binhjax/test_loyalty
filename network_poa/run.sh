HOSTNAME=$(hostname)

NODENAME="node1";
P2PPORT=30301;
RPCPORT=8501;
WSPORT=8541;
ACCOUNT="0xffbcd481c1330e180879b4d2b9b50642eea43c02"

case "$HOSTNAME" in
    "blockchain-2")
        NODENAME="node2";
        P2PPORT=30302;
        RPCPORT=8502;
        WSPORT=8542;
        ACCOUNT="0x2e5b167f68f04918d75f5a6f577a6ea6320225c0"
      ;;
    "blockchain-3")
        NODENAME="node3";
        P2PPORT=30303;
        RPCPORT=8503;
        WSPORT=8543;
        ACCOUNT="0x8f406623e619be85e02b8bb6e4f4ed5c24816e6d"
    ;;
    "block-chain5")
        NODENAME="node4";
        P2PPORT=30304;
        RPCPORT=8504;
        WSPORT=8544;
        ACCOUNT="0x10fff1170de86262d2f65cb81436e40f6c579b44"
    ;;
esac
# nohup
geth \
--port $P2PPORT --rpcport $RPCPORT \
--syncmode full --datadir $NODENAME/datadir \
--ws --wsaddr 0.0.0.0 --wsport $WSPORT --wsorigins="*" \
--rpc --rpcaddr 0.0.0.0   \
--rpcapi 'personal,db,eth,net,web3,txpool,miner,network,debug' \
--networkid 1112 --gasprice 1000   --targetgaslimit '900000000000000000' \
--rpccorsdomain "*" \
--mine  \
--minerthreads 30  \
--unlock "$ACCOUNT" \
--etherbase "$ACCOUNT" \
--nat none \
--password passfile  console
# 2> geth.log &
