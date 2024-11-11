type EutraLocation struct {
	 Tai	Tai	`json:"tai"`
	 IgnoreTai	*bool	`json:"ignoreTai,omitempty"`
	 Ecgi	Ecgi	`json:"ecgi"`
	 IgnoreEcgi	*bool	`json:"ignoreEcgi,omitempty"`
	 GeographicalInformation	string	`json:"geographicalInformation,omitempty"`
	 GlobalNgenbId	*GlobalRanNodeId	`json:"globalNgenbId,omitempty"`
	 GlobalENbId	*GlobalRanNodeId	`json:"globalENbId,omitempty"`
	 AgeOfLocationInformation	*int	`json:"ageOfLocationInformation,omitempty"`
	 UeLocationTimestamp	string	`json:"ueLocationTimestamp,omitempty"`
	 GeodeticInformation	string	`json:"geodeticInformation,omitempty"`
}
