package models

type N2InformationNotification struct {
	N2NotifySubscriptionId string           `json:"n2NotifySubscriptionId"`
	LcsCorrelationId       string           `json:"lcsCorrelationId,omitempty"`
	SmfChangeInfoList      []SmfChangeInfo  `json:"smfChangeInfoList,omitempty"`
	RanNodeId              *GlobalRanNodeId `json:"ranNodeId,omitempty"`
	InitialAmfName         string           `json:"initialAmfName,omitempty"`
	AnN2IPv4Addr           string           `json:"anN2IPv4Addr,omitempty"`
	AnN2IPv6Addr           string           `json:"anN2IPv6Addr,omitempty"`
	Guami                  *Guami           `json:"guami,omitempty"`
	N2InfoContainer        *N2InfoContainer `json:"n2InfoContainer,omitempty"`
	ToReleaseSessionList   []int            `json:"toReleaseSessionList,omitempty"`
	NotifyReason           string           `json:"notifyReason,omitempty"`
	NotifySourceNgRan      *bool            `json:"notifySourceNgRan,omitempty"`
}
