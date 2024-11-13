package models

type EventNotification struct {
	SvcExps            []ServiceExperienceInfo        `json:"svcExps,omitempty"`
	QosSustainInfos    []QosSustainabilityInfo        `json:"qosSustainInfos,omitempty"`
	UserDataCongInfos  []UserDataCongestionInfo       `json:"userDataCongInfos,omitempty"`
	DnPerfInfos        []DnPerfInfo                   `json:"dnPerfInfos,omitempty"`
	DisperInfos        []DispersionInfo               `json:"disperInfos,omitempty"`
	FailNotifyCode     NwdafFailureCode               `json:"failNotifyCode,omitempty"`
	AnaMetaInfo        *AnalyticsMetadataInfo         `json:"anaMetaInfo,omitempty"`
	TimeStampGen       string                         `json:"timeStampGen,omitempty"`
	NwPerfs            []NetworkPerfInfo              `json:"nwPerfs,omitempty"`
	RedTransInfos      []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`
	WlanInfos          []WlanPerformanceInfo          `json:"wlanInfos,omitempty"`
	SmccExps           []SmcceInfo                    `json:"smccExps,omitempty"`
	Event              NwdafEvent                     `json:"event"`
	Expiry             string                         `json:"expiry,omitempty"`
	NsiLoadLevelInfos  []NsiLoadLevelInfo             `json:"nsiLoadLevelInfos,omitempty"`
	UeComms            []UeCommunication              `json:"ueComms,omitempty"`
	UeMobs             []UeMobility                   `json:"ueMobs,omitempty"`
	Start              string                         `json:"start,omitempty"`
	NfLoadLevelInfos   []NfLoadLevelInformation       `json:"nfLoadLevelInfos,omitempty"`
	AbnorBehavrs       []AbnormalBehaviour            `json:"abnorBehavrs,omitempty"`
	RvWaitTime         *int                           `json:"rvWaitTime,omitempty"`
	SliceLoadLevelInfo *SliceLoadLevelInformation     `json:"sliceLoadLevelInfo,omitempty"`
}
