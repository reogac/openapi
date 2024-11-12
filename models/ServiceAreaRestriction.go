package models

type ServiceAreaRestriction struct {
	Areas                         []Area `json:"areas,omitempty"`
	MaxNumOfTAs                   *int   `json:"maxNumOfTAs,omitempty"`
	MaxNumOfTAsForNotAllowedAreas *int   `json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
	RestrictionType               string `json:"restrictionType,omitempty"`
}
