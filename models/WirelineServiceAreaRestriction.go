package models

type WirelineServiceAreaRestriction struct {
	Areas           []WirelineArea  `json:"areas,omitempty"`
	RestrictionType RestrictionType `json:"restrictionType,omitempty"`
}
