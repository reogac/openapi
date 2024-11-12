package models

type EventNotification struct {
	DnPerfInfos        []DnPerfInfo                   `json:"dnPerfInfos,omitempty"`
	RedTransInfos      []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`
	Start              string                         `json:"start,omitempty"`
	AbnorBehavrs       []AbnormalBehaviour            `json:"abnorBehavrs,omitempty"`
	NfLoadLevelInfos   []NfLoadLevelInformation       `json:"nfLoadLevelInfos,omitempty"`
	SliceLoadLevelInfo *SliceLoadLevelInformation     `json:"sliceLoadLevelInfo,omitempty"`
	SvcExps            []ServiceExperienceInfo        `json:"svcExps,omitempty"`
	QosSustainInfos    []QosSustainabilityInfo        `json:"qosSustainInfos,omitempty"`
	UeComms            []UeCommunication              `json:"ueComms,omitempty"`
	NwPerfs            []NetworkPerfInfo              `json:"nwPerfs,omitempty"`
	RvWaitTime         *int                           `json:"rvWaitTime,omitempty"`
	AnaMetaInfo        *AnalyticsMetadataInfo         `json:"anaMetaInfo,omitempty"`
	DisperInfos        []DispersionInfo               `json:"disperInfos,omitempty"`
	SmccExps           []SmcceInfo                    `json:"smccExps,omitempty"`
	UeMobs             []UeMobility                   `json:"ueMobs,omitempty"`
	WlanInfos          []WlanPerformanceInfo          `json:"wlanInfos,omitempty"`
	Event              NwdafEvent                     `json:"event"`
	NsiLoadLevelInfos  []NsiLoadLevelInfo             `json:"nsiLoadLevelInfos,omitempty"`
	FailNotifyCode     NwdafFailureCode               `json:"failNotifyCode,omitempty"`
	UserDataCongInfos  []UserDataCongestionInfo       `json:"userDataCongInfos,omitempty"`
	Expiry             string                         `json:"expiry,omitempty"`
	TimeStampGen       string                         `json:"timeStampGen,omitempty"`
}
