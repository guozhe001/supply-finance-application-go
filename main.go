package main

import (
	"encoding/base64"
	"github.com/guozhe001/supply-finance-application-go/org"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"log"
	"os"
)

func main() {
	//practiceGateWay()
	AccessChannel()
}

func practiceGateWay() {
	log.Println("============ application-golang starts ============")
	setenv()
	conf := []org.ConfInterface{
		org.ConfigCore{},
		org.ConfigF{},
		org.ConfigS1{},
		org.ConfigS2{},
	}
	submitTransaction(getAppAndContract(conf[0]))
	log.Println("============ application-golang ends ============")
}

// 供应链金融项目的所有四个节点
var peers = []string{"peer0.core.supply.com:8051", "peer0.f1.supply.com:8053", "peer0.s1.supply.com:8055", "peer0.s2.supply.com:8151"}

func submitTransaction(app application, contract *gateway.Contract) {
	// 5.Submit the transaction to the network
	//result, err := app.SubmitTransaction(contract, "issue", "MagnetoCorp", "00002", "2020-05-31", "2020-11-30", "5000000")
	// 使用transient数据
	//transient := make(map[string][]byte)
	//privateAsset := "{\"objectType\":\"asset_properties\",\"assetID\":\"application001\",\"issuer\":\"GylCoreOrg1MSP\",\"amount\":1000,\"createDate\":\"2020-01-11T06:57:06.963617Z\",\"endDate\":\"2021-07-11T06:57:06.963617Z\",\"salt\":\"224cba6c547aecc76ab6acfac41d12dfd96e7165\"}"
	//log.Println(privateAsset)
	//base64Asset := base64.StdEncoding.EncodeToString([]byte(privateAsset))
	//log.Println(base64Asset)
	//transient["asset_properties"] = []byte(base64Asset)
	//result, err := app.SubmitTransaction(contract, peers[0:1], transient, "CreateAsset", "application001", "hello application")
	// 查询数据
	result, err := app.SubmitTransaction(contract, []string{}, nil, "ReadAsset", "application001")
	// 查询私有数据库
	//result, err := app.SubmitTransaction(contract, []string{}, nil, "GetAssetPrivateProperties", "application001")
	if err != nil {
		panic(err)
	}
	decodeString, err := base64.StdEncoding.DecodeString(string(result))
	log.Println(string(decodeString))
	// 6.Process the response
	app.ProcessResponse(result)
}

// 获取应用程序接口和合约
func getAppAndContract(c org.ConfInterface) (application, *gateway.Contract) {
	var app application = Application{conf: c}
	// 1.Select an identity from a wallet
	wallet, err := app.GetWallet()
	if err != nil {
		panic(err)
	}
	// 2.Connect to a gateway
	gw, err := app.ConnectGateway(wallet)
	if err != nil {
		panic(err)
	}
	defer gw.Close()
	// 3.Access PaperNet network
	network := app.GetNetWork(gw)
	// 4.Get addressability to commercial paper contract
	contract := app.GetContract(network)
	return app, contract
}

// 设置环境变量，使用本地环境发现网络
func setenv() {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}
}

type application interface {
	// 1.Select an identity from a wallet
	GetWallet() (*gateway.Wallet, error)
	// 2.Connect to a gateway
	ConnectGateway(*gateway.Wallet) (*gateway.Gateway, error)
	// 3.Access PaperNet network
	GetNetWork(*gateway.Gateway) *gateway.Network
	// 4.Get addressability to commercial paper contract
	GetContract(*gateway.Network) *gateway.Contract
	// 5.Submit the transaction to the network
	SubmitTransaction(contract *gateway.Contract, peers []string, transient map[string][]byte, functionName string, args ...string) ([]byte, error)
	// 6.Process the response
	ProcessResponse([]byte)
}
