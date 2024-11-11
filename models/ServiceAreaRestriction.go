type ServiceAreaRestriction struct {
	 MaxNumOfTAsForNotAllowedAreas	*int	`json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
	 RestrictionType	*string	`json:"restrictionType,omitempty"`
	 Areas	[]Area	`json:"areas,omitempty"`
	 MaxNumOfTAs	*int	`json:"maxNumOfTAs,omitempty"`
}
