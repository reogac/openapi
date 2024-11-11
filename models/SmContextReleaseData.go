package models
type SmContextReleaseData struct {
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
	 FiveGMmCauseValue	*int	`json:"5gMmCauseValue,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
	 AddUeLocation	*UserLocation	`json:"addUeLocation,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 UeLocation	*UserLocation	`json:"ueLocation,omitempty"`
	 VsmfReleaseOnly	*bool	`json:"vsmfReleaseOnly,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 N2SmInfoType	string	`json:"n2SmInfoType,omitempty"`
	 IsmfReleaseOnly	*bool	`json:"ismfReleaseOnly,omitempty"`
}
