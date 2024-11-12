package models

type SmfSelectionType string

// Define constant values for SmfSelectionType
const (
	SMFSELECTIONTYPE_CURRENT_PDU_SESSION SmfSelectionType = "CURRENT_PDU_SESSION"
	SMFSELECTIONTYPE_NEXT_PDU_SESSION    SmfSelectionType = "NEXT_PDU_SESSION"
)
