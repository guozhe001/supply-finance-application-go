package main

import (
	"encoding/hex"
	"github.com/guozhe001/supply-finance-application-go/org"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"log"
)

func LedgerPractice() {
	core := org.ConfigCore{}
	client, err := ledger.New(GetChannelContext(core))
	if err != nil {
		log.Panic(err)
	}
	// 根据区块高度查询区块
	block, err := client.QueryBlock(10)
	if err != nil {
		log.Println("QueryBlock", err)
	}
	hash := block.GetHeader().GetDataHash()
	log.Println(hex.EncodeToString(hash))
	log.Println(block)
	QueryBlockByHash(client, hash)

}

func QueryBlockByHash(client *ledger.Client, hash []byte) {
	block, err := client.QueryBlockByHash(hash)
	if err != nil {
		log.Println(err)
	}
	for _, data := range block.GetData().GetData() {
		log.Println("%s", string(data))
	}

}
