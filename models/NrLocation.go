package models
type NrLocation struct {
	 Tai	Tai	`json:"tai"`
	 Ncgi	Ncgi	`json:"ncgi"`
	 IgnoreNcgi	*bool	`json:"ignoreNcgi,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
	 GlobalGnbId	*GlobalRanNodeId	`json:"globalGnbId,omitempty"`
}