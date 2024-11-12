package models

type RetrievedData struct {
	AfCoordinationInfo  *AfCoordinationInfo  `json:"afCoordinationInfo,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
}
