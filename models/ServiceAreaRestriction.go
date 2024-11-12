package models

type ServiceAreaRestriction struct {
	MaxNumOfTAs                   *int            `json:"maxNumOfTAs,omitempty"`
	MaxNumOfTAsForNotAllowedAreas *int            `json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
	RestrictionType               RestrictionType `json:"restrictionType,omitempty"`
	Areas                         []Area          `json:"areas,omitempty"`
}
