package models

type TargetDnaiInfo struct {
	SmfSelectionType SmfSelectionType `json:"smfSelectionType"`
	TargetDnai       string           `json:"targetDnai,omitempty"`
}
