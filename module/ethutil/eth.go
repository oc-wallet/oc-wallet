package ethutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

// ListenNewBlock
func ListenNewBlock() {
	// define chain
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
		return
	}
	// listen block
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:

			fmt.Println("========")

			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
			fmt.Println("Transaction")
			for i, trans := range block.Transactions() {
				fmt.Println("trans i", i, trans.Value())
			}
		}
	}
}
