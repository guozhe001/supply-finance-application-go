package org

import "github.com/guozhe001/supply-finance-application-go/constant"

type ConfigOrg2 struct {
}

func (c ConfigOrg2) GetMspId() string {
	return constant.MspIdCore
}

func (c ConfigOrg2) GetUserName() string {
	return constant.UserName
}

func (c ConfigOrg2) GetCredPathSlice() []string {
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

func (c ConfigOrg2) GetOrgDir() string {
	return constant.OrgDirMain
}

func (c ConfigOrg2) GetConnectionFileName() string {
	return constant.ConnectionFileNameCore
}

func (c ConfigOrg2) GetChannelName() string {
	return constant.ChannelName
}

func (c ConfigOrg2) GetChaincodeId() string {
	return constant.ChaincodeId
}

func (c ConfigOrg2) GetContractName() string {
	return constant.ContractNamePractice
}
