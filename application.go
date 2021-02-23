package main

import (
	"fmt"
	"github.com/guozhe001/supply-finance-application-go/constant"
	"github.com/guozhe001/supply-finance-application-go/org"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Application 核心企业应用程序
type Application struct {
	conf org.ConfInterface
}

// 1.Select an identity from a wallet
func (a Application) GetWallet() (*gateway.Wallet, error) {
	log.Println("1.Select an identity from a wallet")
	///msp/
	walletPath := filepath.Join(
		constant.WalletPath,
		//a.conf.GetMspId(),
		"test",
	)
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		return nil, err

	}

	if !wallet.Exists(a.conf.GetUserName()) {
		credPath := filepath.Join(a.conf.GetCredPathSlice()...)
		//err = populateWallet(wallet, credPath)
		err = populateWallet(a, wallet, credPath)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}
	return wallet, err
}

// 2.Connect to a gateway
func (a Application) ConnectGateway(wallet *gateway.Wallet) (*gateway.Gateway, error) {
	log.Println("2.Connect to a gateway")
	ccpPath := filepath.Join(
		a.conf.GetOrgDir(),
		a.conf.GetConnectionFileName(),
	)

	return gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, a.conf.GetUserName()),
	)
}

// 3.Access PaperNet network
func (a Application) GetNetWork(gw *gateway.Gateway) *gateway.Network {
	network, err := gw.GetNetwork(a.conf.GetChannelName())

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("3.Access PaperNet network, channel=%s", network.Name())
	return network
}

// 4.Get addressability to commercial paper contract
func (a Application) GetContract(network *gateway.Network) *gateway.Contract {
	contract := network.GetContractWithName(a.conf.GetChaincodeId(), a.conf.GetContractName())
	log.Printf("4.Get addressability to commercial paper contract, name=%s,", contract.Name())
	return contract
}

// 5.Submit the transaction to the network
func (a Application) SubmitTransaction(contract *gateway.Contract, endorsingPeers []string, transient map[string][]byte, functionName string, args ...string) ([]byte, error) {
	log.Println("5.Submit the transaction to the network", endorsingPeers, functionName, args)
	txn, err := contract.CreateTransaction(
		functionName,
		gateway.WithEndorsingPeers(endorsingPeers...),
		// 如果指定了不需要签名的peer节点，在有临时数据时这个节点也会收到临时数据从而导致隐私泄漏
		gateway.WithTransient(transient),
	)
	if err != nil {
		fmt.Printf("Failed to create transaction: %s\n", err)
		//return nil, err
	}
	return txn.Submit(args...)
	//return txn.Evaluate(args...)
}

// 6.Process the response
func (a Application) ProcessResponse(bytes []byte) {
	log.Println("6.Process the response")
	log.Printf(string(bytes))
}

func populateWallet(a Application, wallet *gateway.Wallet, credPath string) error {
	log.Println("============ Populating wallet ============")

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyPath := filepath.Join(credPath, "keystore", "priv_sk")
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(a.conf.GetMspId(), string(cert), string(key))

	return wallet.Put(a.conf.GetUserName(), identity)
}
