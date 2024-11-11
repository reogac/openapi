package models

type QosFlowUsageReport struct {
	StartTimeStamp string `json:"startTimeStamp"`
	EndTimeStamp   string `json:"endTimeStamp"`
	DownlinkVolume int64  `json:"downlinkVolume"`
	UplinkVolume   int64  `json:"uplinkVolume"`
	Qfi            int    `json:"qfi"`
}
