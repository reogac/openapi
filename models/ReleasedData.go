package models
type ReleasedData struct {
	 SmallDataRateStatus	*SmallDataRateStatus	`json:"smallDataRateStatus,omitempty"`
	 ApnRateStatus	*ApnRateStatus	`json:"apnRateStatus,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
}
