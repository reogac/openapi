package models

type ConditionData struct {
	DeactivationTime string     `json:"deactivationTime,omitempty"`
	AccessType       AccessType `json:"accessType,omitempty"`
	RatType          RatType    `json:"ratType,omitempty"`
	CondId           string     `json:"condId"`
	ActivationTime   string     `json:"activationTime,omitempty"`
}
