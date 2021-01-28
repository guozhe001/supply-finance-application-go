package org

import "github.com/guozhe001/supply-finance-application-go/constant"

type ConfigF struct {
}

func (c ConfigF) GetMspId() string {
	return constant.MspIdF
}

func (c ConfigF) GetUserName() string {
	return constant.UserName
}

func (c ConfigF) GetCredPathSlice() []string {
	return []string{"/Users",
		"apple",
		"code",
		"open-source",
		"blockchain",
		"hyperledger",
		"supply-finance",
		"organizations",
		"peerOrganizations",
		"f1.supply.com",
		"users",
		"User1@f1.supply.com",
		"msp"}
}

func (c ConfigF) GetOrgDir() string {
	return constant.OrgDirMain
}

func (c ConfigF) GetConnectionFileName() string {
	return constant.ConnectionFileNameF
}

func (c ConfigF) GetChannelName() string {
	return constant.ChannelName
}

func (c ConfigF) GetChaincodeId() string {
	return constant.ChaincodeId
}

func (c ConfigF) GetContractName() string {
	return constant.ContractName
}
