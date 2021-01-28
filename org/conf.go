package org

type ConfInterface interface {
	GetMspId() string
	GetUserName() string
	GetCredPathSlice() []string
	GetOrgDir() string
	GetConnectionFileName() string
	GetChannelName() string
	GetChaincodeId() string
	GetContractName() string
}
