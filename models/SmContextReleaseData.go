type SmContextReleaseData struct {
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
	 5gMmCauseValue	*int	`json:"5gMmCauseValue,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 IsmfReleaseOnly	*bool	`json:"ismfReleaseOnly,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 UeLocation	*UserLocation	`json:"ueLocation,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
	 AddUeLocation	*UserLocation	`json:"addUeLocation,omitempty"`
	 VsmfReleaseOnly	*bool	`json:"vsmfReleaseOnly,omitempty"`
	 N2SmInfoType	string	`json:"n2SmInfoType,omitempty"`
}
