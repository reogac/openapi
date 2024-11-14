package models
type ReleaseData struct {
	 UeLocation	*UserLocation	`json:"ueLocation,omitempty"`
	 AddUeLocation	*UserLocation	`json:"addUeLocation,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 Cause	Cause	`json:"cause,omitempty"`
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
	 SecondaryRatUsageReport	[]SecondaryRatUsageReport	`json:"secondaryRatUsageReport,omitempty"`
	 SecondaryRatUsageInfo	[]SecondaryRatUsageInfo	`json:"secondaryRatUsageInfo,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 FiveGMmCauseValue	*int	`json:"5gMmCauseValue,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
}
