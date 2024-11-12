package models

type AdditionalAccessInfo struct {
	AccessType AccessType `json:"accessType"`
	RatType    RatType    `json:"ratType,omitempty"`
}
