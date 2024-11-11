package models

type VolumeTimedReport struct {
	EndTimeStamp   string `json:"endTimeStamp"`
	DownlinkVolume int64  `json:"downlinkVolume"`
	UplinkVolume   int64  `json:"uplinkVolume"`
	StartTimeStamp string `json:"startTimeStamp"`
}
