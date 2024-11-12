package models

type SmContextReleasedData struct {
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
}
