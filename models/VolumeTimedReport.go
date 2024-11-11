package models
type VolumeTimedReport struct {
	 UplinkVolume	int64	`json:"uplinkVolume"`
	 StartTimeStamp	string	`json:"startTimeStamp"`
	 EndTimeStamp	string	`json:"endTimeStamp"`
	 DownlinkVolume	int64	`json:"downlinkVolume"`
}
