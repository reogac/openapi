package models
type QosFlowUsageReport struct {
	 UplinkVolume	int64	`json:"uplinkVolume"`
	 Qfi	int	`json:"qfi"`
	 StartTimeStamp	string	`json:"startTimeStamp"`
	 EndTimeStamp	string	`json:"endTimeStamp"`
	 DownlinkVolume	int64	`json:"downlinkVolume"`
}
