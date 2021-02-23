package org

import "github.com/guozhe001/supply-finance-application-go/constant"

type ConfigSupply struct {
}

func (c ConfigSupply) GetMspId() string {
	return constant.MspIdCore
}

func (c ConfigSupply) GetUserName() string {
	return constant.UserName
}

func (c ConfigSupply) GetCredPathSlice() []string {
	//return []string{"/Users",
	//	"apple",
	//	"code",
	//	"open-source",
	//	"blockchain",
	//	"hyperledger",
	//	"supply-finance",
	//	"organizations",
	//	"peerOrganizations",
	//	"core.supply.com",
	//	"users",
	//	//"Admin@core.supply.com",
	//	"User1@core.supply.com",
	//	"msp"}
	return []string{
		"/Users",
		"apple",
		"code",
		"open-source",
		"blockchain",
		"hyperledger",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp"}
}

func (c ConfigSupply) GetOrgDir() string {
	return constant.OrgDirMain
}

func (c ConfigSupply) GetConnectionFileName() string {
	return constant.ConnectionFileNameSupply
}

func (c ConfigSupply) GetChannelName() string {
	return "mychannel"
}

func (c ConfigSupply) GetChaincodeId() string {
	return "supply"
}

func (c ConfigSupply) GetContractName() string {
	return ""
}
