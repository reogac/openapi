package models

type EventNotification struct {
	DisperInfos        []DispersionInfo               `json:"disperInfos,omitempty"`
	WlanInfos          []WlanPerformanceInfo          `json:"wlanInfos,omitempty"`
	SmccExps           []SmcceInfo                    `json:"smccExps,omitempty"`
	FailNotifyCode     NwdafFailureCode               `json:"failNotifyCode,omitempty"`
	RvWaitTime         *int                           `json:"rvWaitTime,omitempty"`
	AnaMetaInfo        *AnalyticsMetadataInfo         `json:"anaMetaInfo,omitempty"`
	QosSustainInfos    []QosSustainabilityInfo        `json:"qosSustainInfos,omitempty"`
	SliceLoadLevelInfo *SliceLoadLevelInformation     `json:"sliceLoadLevelInfo,omitempty"`
	SvcExps            []ServiceExperienceInfo        `json:"svcExps,omitempty"`
	UeMobs             []UeMobility                   `json:"ueMobs,omitempty"`
	DnPerfInfos        []DnPerfInfo                   `json:"dnPerfInfos,omitempty"`
	Event              NwdafEvent                     `json:"event"`
	Start              string                         `json:"start,omitempty"`
	NfLoadLevelInfos   []NfLoadLevelInformation       `json:"nfLoadLevelInfos,omitempty"`
	NsiLoadLevelInfos  []NsiLoadLevelInfo             `json:"nsiLoadLevelInfos,omitempty"`
	Expiry             string                         `json:"expiry,omitempty"`
	TimeStampGen       string                         `json:"timeStampGen,omitempty"`
	UeComms            []UeCommunication              `json:"ueComms,omitempty"`
	RedTransInfos      []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`
	UserDataCongInfos  []UserDataCongestionInfo       `json:"userDataCongInfos,omitempty"`
	AbnorBehavrs       []AbnormalBehaviour            `json:"abnorBehavrs,omitempty"`
	NwPerfs            []NetworkPerfInfo              `json:"nwPerfs,omitempty"`
}
