package org

import "github.com/guozhe001/supply-finance/application-go/constant"

type ConfigS1 struct {
}

func (s ConfigS1) GetMspId() string {
	return constant.MspIdS1
}

func (s ConfigS1) GetUserName() string {
	return constant.UserName
}

func (s ConfigS1) GetCredPathSlice() []string {
	return []string{"/Users",
		"apple",
		"code",
		"open-source",
		"blockchain",
		"hyperledger",
		"supply-finance",
		"organizations",
		"peerOrganizations",
		"s1.supply.com",
		"users",
		"User1@s1.supply.com",
		"msp"}
}

func (s ConfigS1) GetOrgDir() string {
	return constant.OrgDirMain
}

func (s ConfigS1) GetConnectionFileName() string {
	return constant.ConnectionFileNameS1
}

func (s ConfigS1) GetChannelName() string {
	return constant.ChannelName
}

func (s ConfigS1) GetChaincodeId() string {
	return constant.ChaincodeId
}

func (s ConfigS1) GetContractName() string {
	return constant.ContractName
}
