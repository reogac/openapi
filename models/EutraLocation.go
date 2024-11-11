package models
type EutraLocation struct {
	 IgnoreTai	*bool	`json:"ignoreTai,omitempty"`
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GlobalNgenbId	*GlobalRanNodeId	`json:"globalNgenbId,omitempty"`
	 GlobalENbId	*GlobalRanNodeId	`json:"globalENbId,omitempty"`
	 Tai	Tai	`json:"tai"`
	 Ecgi	Ecgi	`json:"ecgi"`
	 IgnoreEcgi	*bool	`json:"ignoreEcgi,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
}
