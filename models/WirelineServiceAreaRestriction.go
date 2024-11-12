package models
type WirelineServiceAreaRestriction struct {
	 RestrictionType	RestrictionType	`json:"restrictionType,omitempty"`
	 Areas	[]WirelineArea	`json:"areas,omitempty"`
}
