package models

type QosFlowUsageReport struct {
	EndTimeStamp   string `json:"endTimeStamp"`
	DownlinkVolume int64  `json:"downlinkVolume"`
	UplinkVolume   int64  `json:"uplinkVolume"`
	Qfi            int    `json:"qfi"`
	StartTimeStamp string `json:"startTimeStamp"`
}
