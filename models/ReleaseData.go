type ReleaseData struct {
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
	 UeLocation	*UserLocation	`json:"ueLocation,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
	 SecondaryRatUsageInfo	[]SecondaryRatUsageInfo	`json:"secondaryRatUsageInfo,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 5gMmCauseValue	*int	`json:"5gMmCauseValue,omitempty"`
	 AddUeLocation	*UserLocation	`json:"addUeLocation,omitempty"`
	 SecondaryRatUsageReport	[]SecondaryRatUsageReport	`json:"secondaryRatUsageReport,omitempty"`
}
