package models

type UcPurpose string

// Define constant values for UcPurpose
const (
	UCPURPOSE_ANALYTICS           UcPurpose = "ANALYTICS"
	UCPURPOSE_MODEL_TRAINING      UcPurpose = "MODEL_TRAINING"
	UCPURPOSE_NW_CAP_EXPOSURE     UcPurpose = "NW_CAP_EXPOSURE"
	UCPURPOSE_EDGEAPP_UE_LOCATION UcPurpose = "EDGEAPP_UE_LOCATION"
)
