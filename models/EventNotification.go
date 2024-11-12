package models
type EventNotification struct {
	 SmccExps	[]SmcceInfo	`json:"smccExps,omitempty"`
	 TimeStampGen	string	`json:"timeStampGen,omitempty"`
	 SvcExps	[]ServiceExperienceInfo	`json:"svcExps,omitempty"`
	 UeMobs	[]UeMobility	`json:"ueMobs,omitempty"`
	 AbnorBehavrs	[]AbnormalBehaviour	`json:"abnorBehavrs,omitempty"`
	 NsiLoadLevelInfos	[]NsiLoadLevelInfo	`json:"nsiLoadLevelInfos,omitempty"`
	 DnPerfInfos	[]DnPerfInfo	`json:"dnPerfInfos,omitempty"`
	 Event	NwdafEvent	`json:"event"`
	 FailNotifyCode	NwdafFailureCode	`json:"failNotifyCode,omitempty"`
	 RvWaitTime	*int	`json:"rvWaitTime,omitempty"`
	 AnaMetaInfo	*AnalyticsMetadataInfo	`json:"anaMetaInfo,omitempty"`
	 RedTransInfos	[]RedundantTransmissionExpInfo	`json:"redTransInfos,omitempty"`
	 WlanInfos	[]WlanPerformanceInfo	`json:"wlanInfos,omitempty"`
	 Expiry	string	`json:"expiry,omitempty"`
	 NfLoadLevelInfos	[]NfLoadLevelInformation	`json:"nfLoadLevelInfos,omitempty"`
	 SliceLoadLevelInfo	*SliceLoadLevelInformation	`json:"sliceLoadLevelInfo,omitempty"`
	 UeComms	[]UeCommunication	`json:"ueComms,omitempty"`
	 DisperInfos	[]DispersionInfo	`json:"disperInfos,omitempty"`
	 Start	string	`json:"start,omitempty"`
	 QosSustainInfos	[]QosSustainabilityInfo	`json:"qosSustainInfos,omitempty"`
	 UserDataCongInfos	[]UserDataCongestionInfo	`json:"userDataCongInfos,omitempty"`
	 NwPerfs	[]NetworkPerfInfo	`json:"nwPerfs,omitempty"`
}
