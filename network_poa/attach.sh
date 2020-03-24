HOSTNAME=$(hostname)
NODENAME="node1"
case "$HOSTNAME" in
    "blockchain-2")   NODENAME="node2"
    ;;
    "blockchain-3")   NODENAME="node3"
    ;;
    "block-chain5")   NODENAME="node4"
    ;;
esac
geth attach ipc://app/ethereum/network_eth_poa/$NODENAME/datadir/geth.ipc
