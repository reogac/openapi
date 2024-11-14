package models
type SmContextReleaseData struct {
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
	 FiveGMmCauseValue	*int	`json:"5gMmCauseValue,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
	 VsmfReleaseOnly	*bool	`json:"vsmfReleaseOnly,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 N2SmInfoType	N2SmInfoType	`json:"n2SmInfoType,omitempty"`
	 Cause	Cause	`json:"cause,omitempty"`
	 AddUeLocation	*UserLocation	`json:"addUeLocation,omitempty"`
	 IsmfReleaseOnly	*bool	`json:"ismfReleaseOnly,omitempty"`
	 UeLocation	*UserLocation	`json:"ueLocation,omitempty"`
}
