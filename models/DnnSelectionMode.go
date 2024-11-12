package models

type DnnSelectionMode string

// Define constant values for DnnSelectionMode
const (
	DNNSELECTIONMODE_VERIFIED            DnnSelectionMode = "VERIFIED"
	DNNSELECTIONMODE_UE_DNN_NOT_VERIFIED DnnSelectionMode = "UE_DNN_NOT_VERIFIED"
	DNNSELECTIONMODE_NW_DNN_NOT_VERIFIED DnnSelectionMode = "NW_DNN_NOT_VERIFIED"
)
