package models

type EventNotification struct {
	TimeStampGen       string                         `json:"timeStampGen,omitempty"`
	RvWaitTime         *int                           `json:"rvWaitTime,omitempty"`
	AbnorBehavrs       []AbnormalBehaviour            `json:"abnorBehavrs,omitempty"`
	RedTransInfos      []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`
	WlanInfos          []WlanPerformanceInfo          `json:"wlanInfos,omitempty"`
	AnaMetaInfo        *AnalyticsMetadataInfo         `json:"anaMetaInfo,omitempty"`
	NsiLoadLevelInfos  []NsiLoadLevelInfo             `json:"nsiLoadLevelInfos,omitempty"`
	SmccExps           []SmcceInfo                    `json:"smccExps,omitempty"`
	DnPerfInfos        []DnPerfInfo                   `json:"dnPerfInfos,omitempty"`
	Event              NwdafEvent                     `json:"event"`
	SliceLoadLevelInfo *SliceLoadLevelInformation     `json:"sliceLoadLevelInfo,omitempty"`
	QosSustainInfos    []QosSustainabilityInfo        `json:"qosSustainInfos,omitempty"`
	UeMobs             []UeMobility                   `json:"ueMobs,omitempty"`
	UserDataCongInfos  []UserDataCongestionInfo       `json:"userDataCongInfos,omitempty"`
	NwPerfs            []NetworkPerfInfo              `json:"nwPerfs,omitempty"`
	DisperInfos        []DispersionInfo               `json:"disperInfos,omitempty"`
	Start              string                         `json:"start,omitempty"`
	Expiry             string                         `json:"expiry,omitempty"`
	FailNotifyCode     NwdafFailureCode               `json:"failNotifyCode,omitempty"`
	NfLoadLevelInfos   []NfLoadLevelInformation       `json:"nfLoadLevelInfos,omitempty"`
	SvcExps            []ServiceExperienceInfo        `json:"svcExps,omitempty"`
	UeComms            []UeCommunication              `json:"ueComms,omitempty"`
}
