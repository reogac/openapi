package models

type HssAuthTypeInUri string

// Define constant values for HssAuthTypeInUri
const (
	HSSAUTHTYPEINURI_EPS_AKA       HssAuthTypeInUri = "eps-aka"
	HSSAUTHTYPEINURI_EAP_AKA       HssAuthTypeInUri = "eap-aka"
	HSSAUTHTYPEINURI_EAP_AKA_PRIME HssAuthTypeInUri = "eap-aka-prime"
	HSSAUTHTYPEINURI_IMS_AKA       HssAuthTypeInUri = "ims-aka"
	HSSAUTHTYPEINURI_GBA_AKA       HssAuthTypeInUri = "gba-aka"
)