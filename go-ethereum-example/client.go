package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//如果您没有现有的以太坊客户端，您可以连接到infura网关。Infura管理着一个安全、可靠、可扩展的以太坊[geth和parity]节点，并且在接入以太坊网络时减少了新人的入门门户。
	//client, err := ethclient.Dial("http://localhost:8545")
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
