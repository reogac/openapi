package models

type HssAuthType string

// Define constant values for HssAuthType
const (
	HSSAUTHTYPE_EPS_AKA       HssAuthType = "EPS_AKA"
	HSSAUTHTYPE_EAP_AKA       HssAuthType = "EAP_AKA"
	HSSAUTHTYPE_EAP_AKA_PRIME HssAuthType = "EAP_AKA_PRIME"
	HSSAUTHTYPE_IMS_AKA       HssAuthType = "IMS_AKA"
	HSSAUTHTYPE_GBA_AKA       HssAuthType = "GBA_AKA"
	HSSAUTHTYPE_UMTS_AKA      HssAuthType = "UMTS_AKA"
)