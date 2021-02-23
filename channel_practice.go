package main

import (
	"github.com/guozhe001/supply-finance-application-go/constant"
	"github.com/guozhe001/supply-finance-application-go/org"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel/invoke"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
)

func channelPractice() {
	var orgConfig = org.ConfigCore{}
	cc := GetChannelClient(orgConfig)
	invokeAndPrintLog(Query, cc, orgConfig)
	invokeAndPrintLog(Execute, cc, orgConfig)
	invokeAndPrintLog(InvokeHandler, cc, orgConfig)
	// 注册和注销链码的事件
	//cc.RegisterChaincodeEvent()
	//cc.UnregisterChaincodeEvent()
}

func invokeAndPrintLog(f func(*channel.Client, org.ConfigCore), cc *channel.Client, orgConfig org.ConfigCore) {
	log.Printf("invoke method %s =========================================", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	f(cc, orgConfig)
}

func InvokeHandler(cc *channel.Client, orgConfig org.ConfigCore) {
	// 通过channel.Client的InvokeHandler方法调用链码，可以传入自定义的handler
	handler, err := cc.InvokeHandler(MyHandler{}, channel.Request{ChaincodeID: orgConfig.GetChaincodeId(), Fcn: "ReadAsset", Args: [][]byte{[]byte("application001")}})
	if err != nil {
		log.Printf("failed to Execute chaincode: %s\n", err)
	}
	log.Printf("handler=%#v", handler)
}

type MyHandler struct {
}

func (m MyHandler) Handle(context *invoke.RequestContext, clientContext *invoke.ClientContext) {
	log.Printf("hello i'm MyHandler; Request=%#v Response=%v#", context.Request, context.Response)
}

func Execute(cc *channel.Client, orgConfig org.ConfigCore) {
	// 通过channel.Client执行链码chaincode，Execute方法在执行完链码之后会把交易发给order节点
	response, err := cc.Execute(channel.Request{ChaincodeID: orgConfig.GetChaincodeId(), Fcn: "ReadAsset", Args: [][]byte{[]byte("application001")}})
	if err != nil {
		log.Printf("failed to Execute chaincode: %s\n", err)
	}
	log.Println(string(response.Payload))
}

func Query(cc *channel.Client, orgConfig org.ConfigCore) {
	// 通过channel.Client查询链码chaincode，Query方法只是查询，不会根order节点交互
	response, err := cc.Query(channel.Request{ChaincodeID: orgConfig.GetChaincodeId(), Fcn: "ReadAsset", Args: [][]byte{[]byte("application001")}})
	if err != nil {
		log.Printf("failed to query chaincode: %s\n", err)
	}
	log.Println(string(response.Payload))
}

// GetChannelClient 获取channel.Client
func GetChannelClient(orgConfig org.ConfInterface) *channel.Client {
	cc, err := channel.New(GetChannelContext(orgConfig))
	if err != nil {
		log.Panicf("failed to create channel client: %s", err)
	}
	return cc
}

// 获取通道上下文
func GetChannelContext(orgConfig org.ConfInterface) context.ChannelProvider {
	//rcp := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("GylCoreOrg1MSP"))
	//rc, err := resmgmt.New(rcp)
	return GetSdk(orgConfig).ChannelContext(orgConfig.GetChannelName(), fabsdk.WithUser(constant.Admin))
}

// 获取SDK实例
func GetSdk(orgConfig org.ConfInterface) *fabsdk.FabricSDK {
	sdk, err := fabsdk.New(config.FromFile(filepath.Join(constant.OrgDirMain, orgConfig.GetConnectionFileName())))
	if err != nil {
		log.Panicf("failed to create fabric sdk: %s", err)
	}
	return sdk
}
