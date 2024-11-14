package models
type EventNotification struct {
	 AnaMetaInfo	*AnalyticsMetadataInfo	`json:"anaMetaInfo,omitempty"`
	 NfLoadLevelInfos	[]NfLoadLevelInformation	`json:"nfLoadLevelInfos,omitempty"`
	 NsiLoadLevelInfos	[]NsiLoadLevelInfo	`json:"nsiLoadLevelInfos,omitempty"`
	 SvcExps	[]ServiceExperienceInfo	`json:"svcExps,omitempty"`
	 UeComms	[]UeCommunication	`json:"ueComms,omitempty"`
	 DisperInfos	[]DispersionInfo	`json:"disperInfos,omitempty"`
	 SmccExps	[]SmcceInfo	`json:"smccExps,omitempty"`
	 TimeStampGen	string	`json:"timeStampGen,omitempty"`
	 FailNotifyCode	NwdafFailureCode	`json:"failNotifyCode,omitempty"`
	 SliceLoadLevelInfo	*SliceLoadLevelInformation	`json:"sliceLoadLevelInfo,omitempty"`
	 NwPerfs	[]NetworkPerfInfo	`json:"nwPerfs,omitempty"`
	 DnPerfInfos	[]DnPerfInfo	`json:"dnPerfInfos,omitempty"`
	 RedTransInfos	[]RedundantTransmissionExpInfo	`json:"redTransInfos,omitempty"`
	 Event	NwdafEvent	`json:"event"`
	 RvWaitTime	*int	`json:"rvWaitTime,omitempty"`
	 QosSustainInfos	[]QosSustainabilityInfo	`json:"qosSustainInfos,omitempty"`
	 UeMobs	[]UeMobility	`json:"ueMobs,omitempty"`
	 WlanInfos	[]WlanPerformanceInfo	`json:"wlanInfos,omitempty"`
	 Start	string	`json:"start,omitempty"`
	 UserDataCongInfos	[]UserDataCongestionInfo	`json:"userDataCongInfos,omitempty"`
	 AbnorBehavrs	[]AbnormalBehaviour	`json:"abnorBehavrs,omitempty"`
	 Expiry	string	`json:"expiry,omitempty"`
}
