package models

type RequestedRuleDataType string

// Define constant values for RequestedRuleDataType
const (
	REQUESTEDRULEDATATYPE_CH_ID         RequestedRuleDataType = "CH_ID"
	REQUESTEDRULEDATATYPE_MS_TIME_ZONE  RequestedRuleDataType = "MS_TIME_ZONE"
	REQUESTEDRULEDATATYPE_USER_LOC_INFO RequestedRuleDataType = "USER_LOC_INFO"
	REQUESTEDRULEDATATYPE_RES_RELEASE   RequestedRuleDataType = "RES_RELEASE"
	REQUESTEDRULEDATATYPE_SUCC_RES_ALLO RequestedRuleDataType = "SUCC_RES_ALLO"
	REQUESTEDRULEDATATYPE_EPS_FALLBACK  RequestedRuleDataType = "EPS_FALLBACK"
)