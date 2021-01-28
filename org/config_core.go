package org

import "github.com/guozhe001/supply-finance-application-go/constant"

type ConfigCore struct {
}

func (c ConfigCore) GetMspId() string {
	return constant.MspIdCore
}

func (c ConfigCore) GetUserName() string {
	return constant.UserName
}

func (c ConfigCore) GetCredPathSlice() []string {
	return []string{"/Users",
		"apple",
		"code",
		"open-source",
		"blockchain",
		"hyperledger",
		"supply-finance",
		"organizations",
		"peerOrganizations",
		"core.supply.com",
		"users",
		//"Admin@core.supply.com",
		"User1@core.supply.com",
		"msp"}
}

func (c ConfigCore) GetOrgDir() string {
	return constant.OrgDirMain
}

func (c ConfigCore) GetConnectionFileName() string {
	return constant.ConnectionFileNameCore
}

func (c ConfigCore) GetChannelName() string {
	return constant.ChannelName
}

func (c ConfigCore) GetChaincodeId() string {
	return constant.ChaincodeId
}

func (c ConfigCore) GetContractName() string {
	return constant.ContractName
}
