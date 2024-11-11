package models
type ServiceAreaRestriction struct {
	 RestrictionType	string	`json:"restrictionType,omitempty"`
	 Areas	[]Area	`json:"areas,omitempty"`
	 MaxNumOfTAs	*int	`json:"maxNumOfTAs,omitempty"`
	 MaxNumOfTAsForNotAllowedAreas	*int	`json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
}
