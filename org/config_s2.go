package org

import "github.com/guozhe001/supply-finance-application-go/constant"

type ConfigS2 struct {
}

func (c ConfigS2) GetMspId() string {
	return constant.MspIdS2
}

func (c ConfigS2) GetUserName() string {
	return constant.UserName
}

func (c ConfigS2) GetCredPathSlice() []string {
	return []string{"/Users",
		"apple",
		"code",
		"open-source",
		"blockchain",
		"hyperledger",
		"supply-finance",
		"organizations",
		"peerOrganizations",
		"s2.supply.com",
		"users",
		"User1@s2.supply.com",
		"msp"}
}

func (c ConfigS2) GetOrgDir() string {
	return constant.OrgDirMain
}

func (c ConfigS2) GetConnectionFileName() string {
	return constant.ConnectionFileNameS2
}

func (c ConfigS2) GetChannelName() string {
	return constant.ChannelName
}

func (c ConfigS2) GetChaincodeId() string {
	return constant.ChaincodeId
}

func (c ConfigS2) GetContractName() string {
	return constant.ContractName
}
