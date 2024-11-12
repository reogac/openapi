package models

type Lpi struct {
	LocationPrivacyInd LocationPrivacyInd `json:"locationPrivacyInd"`
	ValidTimePeriod    *ValidTimePeriod   `json:"validTimePeriod,omitempty"`
}
