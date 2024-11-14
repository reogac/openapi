package models
type EventNotification struct {
	 UserDataCongInfos	[]UserDataCongestionInfo	`json:"userDataCongInfos,omitempty"`
	 AbnorBehavrs	[]AbnormalBehaviour	`json:"abnorBehavrs,omitempty"`
	 FailNotifyCode	NwdafFailureCode	`json:"failNotifyCode,omitempty"`
	 AnaMetaInfo	*AnalyticsMetadataInfo	`json:"anaMetaInfo,omitempty"`
	 UeMobs	[]UeMobility	`json:"ueMobs,omitempty"`
	 DnPerfInfos	[]DnPerfInfo	`json:"dnPerfInfos,omitempty"`
	 RedTransInfos	[]RedundantTransmissionExpInfo	`json:"redTransInfos,omitempty"`
	 TimeStampGen	string	`json:"timeStampGen,omitempty"`
	 SliceLoadLevelInfo	*SliceLoadLevelInformation	`json:"sliceLoadLevelInfo,omitempty"`
	 QosSustainInfos	[]QosSustainabilityInfo	`json:"qosSustainInfos,omitempty"`
	 SvcExps	[]ServiceExperienceInfo	`json:"svcExps,omitempty"`
	 NwPerfs	[]NetworkPerfInfo	`json:"nwPerfs,omitempty"`
	 Event	NwdafEvent	`json:"event"`
	 Expiry	string	`json:"expiry,omitempty"`
	 RvWaitTime	*int	`json:"rvWaitTime,omitempty"`
	 UeComms	[]UeCommunication	`json:"ueComms,omitempty"`
	 DisperInfos	[]DispersionInfo	`json:"disperInfos,omitempty"`
	 WlanInfos	[]WlanPerformanceInfo	`json:"wlanInfos,omitempty"`
	 SmccExps	[]SmcceInfo	`json:"smccExps,omitempty"`
	 Start	string	`json:"start,omitempty"`
	 NfLoadLevelInfos	[]NfLoadLevelInformation	`json:"nfLoadLevelInfos,omitempty"`
	 NsiLoadLevelInfos	[]NsiLoadLevelInfo	`json:"nsiLoadLevelInfos,omitempty"`
}
