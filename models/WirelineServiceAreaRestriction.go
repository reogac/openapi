package models
type WirelineServiceAreaRestriction struct {
	 RestrictionType	string	`json:"restrictionType,omitempty"`
	 Areas	[]WirelineArea	`json:"areas,omitempty"`
}
