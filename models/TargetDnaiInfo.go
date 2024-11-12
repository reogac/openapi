package models

type TargetDnaiInfo struct {
	TargetDnai       string           `json:"targetDnai,omitempty"`
	SmfSelectionType SmfSelectionType `json:"smfSelectionType"`
}
