package model

type P map[string]interface{}

type Dict map[interface{}]interface{}

type PolicyType int

const (
	POLICY_TYPE_ACCESS    PolicyType = 0
	POLICY_TYPE_DATAMASK  PolicyType = 1
	POLICY_TYPE_ROWFILTER PolicyType = 2
)
