package models
type QosFlowUsageReport struct {
	 Qfi	int	`json:"qfi"`
	 StartTimeStamp	string	`json:"startTimeStamp"`
	 EndTimeStamp	string	`json:"endTimeStamp"`
	 DownlinkVolume	int64	`json:"downlinkVolume"`
	 UplinkVolume	int64	`json:"uplinkVolume"`
}
