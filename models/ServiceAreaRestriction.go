package models
type ServiceAreaRestriction struct {
	 RestrictionType	RestrictionType	`json:"restrictionType,omitempty"`
	 Areas	[]Area	`json:"areas,omitempty"`
	 MaxNumOfTAs	*int	`json:"maxNumOfTAs,omitempty"`
	 MaxNumOfTAsForNotAllowedAreas	*int	`json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
}
