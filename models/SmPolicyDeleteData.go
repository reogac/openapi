package models
type SmPolicyDeleteData struct {
	 RanNasRelCauses	[]RanNasRelCause	`json:"ranNasRelCauses,omitempty"`
	 AccuUsageReports	[]AccuUsageReport	`json:"accuUsageReports,omitempty"`
	 PduSessRelCause	PduSessionRelCause	`json:"pduSessRelCause,omitempty"`
	 UserLocationInfo	*UserLocation	`json:"userLocationInfo,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
	 ServingNetwork	*PlmnIdNid	`json:"servingNetwork,omitempty"`
	 UserLocationInfoTime	string	`json:"userLocationInfoTime,omitempty"`
}
