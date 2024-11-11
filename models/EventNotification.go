package models

type EventNotification struct {
	Event              string                         `json:"event"`
	Start              string                         `json:"start,omitempty"`
	NfLoadLevelInfos   []NfLoadLevelInformation       `json:"nfLoadLevelInfos,omitempty"`
	SmccExps           []SmcceInfo                    `json:"smccExps,omitempty"`
	TimeStampGen       string                         `json:"timeStampGen,omitempty"`
	FailNotifyCode     string                         `json:"failNotifyCode,omitempty"`
	QosSustainInfos    []QosSustainabilityInfo        `json:"qosSustainInfos,omitempty"`
	UeComms            []UeCommunication              `json:"ueComms,omitempty"`
	DnPerfInfos        []DnPerfInfo                   `json:"dnPerfInfos,omitempty"`
	WlanInfos          []WlanPerformanceInfo          `json:"wlanInfos,omitempty"`
	NwPerfs            []NetworkPerfInfo              `json:"nwPerfs,omitempty"`
	DisperInfos        []DispersionInfo               `json:"disperInfos,omitempty"`
	Expiry             string                         `json:"expiry,omitempty"`
	RvWaitTime         *int                           `json:"rvWaitTime,omitempty"`
	AnaMetaInfo        *AnalyticsMetadataInfo         `json:"anaMetaInfo,omitempty"`
	NsiLoadLevelInfos  []NsiLoadLevelInfo             `json:"nsiLoadLevelInfos,omitempty"`
	SliceLoadLevelInfo *SliceLoadLevelInformation     `json:"sliceLoadLevelInfo,omitempty"`
	UserDataCongInfos  []UserDataCongestionInfo       `json:"userDataCongInfos,omitempty"`
	SvcExps            []ServiceExperienceInfo        `json:"svcExps,omitempty"`
	UeMobs             []UeMobility                   `json:"ueMobs,omitempty"`
	AbnorBehavrs       []AbnormalBehaviour            `json:"abnorBehavrs,omitempty"`
	RedTransInfos      []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`
}
