package models

type HssAvType string

// Define constant values for HssAvType
const (
	HSSAVTYPE_EPS_AKA  HssAvType = "EPS_AKA"
	HSSAVTYPE_EAP_AKA  HssAvType = "EAP_AKA"
	HSSAVTYPE_IMS_AKA  HssAvType = "IMS_AKA"
	HSSAVTYPE_GBA_AKA  HssAvType = "GBA_AKA"
	HSSAVTYPE_UMTS_AKA HssAvType = "UMTS_AKA"
)