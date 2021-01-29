package main

import (
	"github.com/guozhe001/supply-finance-application-go/org"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"log"
)

func eventPractice() {
	core := org.ConfigCore{}
	context := GetChannelContext(core)
	client, err := event.New(context, event.WithBlockEvents())
	if err != nil {
		log.Panic(err)
	}
	//
	//chaincodeEvent, events, err := client.RegisterChaincodeEvent(core.GetChaincodeId(), "*")
	// 注册区块事件
	blockEvent, _, err := client.RegisterBlockEvent(MyFilter)
	// 注销事件
	client.Unregister(blockEvent)
	// 注册
	filteredBlockEvent, _, err := client.RegisterFilteredBlockEvent()
	client.Unregister(filteredBlockEvent)
	// 注册交易事件
	statusEvent, _, err := client.RegisterTxStatusEvent("txId")
	client.Unregister(statusEvent)
}

func MyFilter(block *common.Block) bool {
	if block.GetHeader().Number > 1 {
		return true
	}
	return false
}
