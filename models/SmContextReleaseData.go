package models
type SmContextReleaseData struct {
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
	 FiveGMmCauseValue	*int	`json:"5gMmCauseValue,omitempty"`
	 UeLocation	*UserLocation	`json:"ueLocation,omitempty"`
	 VsmfReleaseOnly	*bool	`json:"vsmfReleaseOnly,omitempty"`
	 Cause	Cause	`json:"cause,omitempty"`
	 AddUeLocation	*UserLocation	`json:"addUeLocation,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 N2SmInfoType	N2SmInfoType	`json:"n2SmInfoType,omitempty"`
	 IsmfReleaseOnly	*bool	`json:"ismfReleaseOnly,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
}
