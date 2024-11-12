package models

type EventNotification struct {
	QosSustainInfos    []QosSustainabilityInfo        `json:"qosSustainInfos,omitempty"`
	NwPerfs            []NetworkPerfInfo              `json:"nwPerfs,omitempty"`
	FailNotifyCode     string                         `json:"failNotifyCode,omitempty"`
	NfLoadLevelInfos   []NfLoadLevelInformation       `json:"nfLoadLevelInfos,omitempty"`
	NsiLoadLevelInfos  []NsiLoadLevelInfo             `json:"nsiLoadLevelInfos,omitempty"`
	RvWaitTime         *int                           `json:"rvWaitTime,omitempty"`
	AnaMetaInfo        *AnalyticsMetadataInfo         `json:"anaMetaInfo,omitempty"`
	SvcExps            []ServiceExperienceInfo        `json:"svcExps,omitempty"`
	DisperInfos        []DispersionInfo               `json:"disperInfos,omitempty"`
	WlanInfos          []WlanPerformanceInfo          `json:"wlanInfos,omitempty"`
	Event              string                         `json:"event"`
	Expiry             string                         `json:"expiry,omitempty"`
	TimeStampGen       string                         `json:"timeStampGen,omitempty"`
	SmccExps           []SmcceInfo                    `json:"smccExps,omitempty"`
	UserDataCongInfos  []UserDataCongestionInfo       `json:"userDataCongInfos,omitempty"`
	AbnorBehavrs       []AbnormalBehaviour            `json:"abnorBehavrs,omitempty"`
	Start              string                         `json:"start,omitempty"`
	UeComms            []UeCommunication              `json:"ueComms,omitempty"`
	UeMobs             []UeMobility                   `json:"ueMobs,omitempty"`
	SliceLoadLevelInfo *SliceLoadLevelInformation     `json:"sliceLoadLevelInfo,omitempty"`
	DnPerfInfos        []DnPerfInfo                   `json:"dnPerfInfos,omitempty"`
	RedTransInfos      []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`
}
