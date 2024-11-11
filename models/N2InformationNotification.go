package models

type N2InformationNotification struct {
	NotifySourceNgRan      *bool            `json:"notifySourceNgRan,omitempty"`
	N2NotifySubscriptionId string           `json:"n2NotifySubscriptionId"`
	N2InfoContainer        *N2InfoContainer `json:"n2InfoContainer,omitempty"`
	LcsCorrelationId       string           `json:"lcsCorrelationId,omitempty"`
	NotifyReason           string           `json:"notifyReason,omitempty"`
	InitialAmfName         string           `json:"initialAmfName,omitempty"`
	AnN2IPv6Addr           string           `json:"anN2IPv6Addr,omitempty"`
	Guami                  *Guami           `json:"guami,omitempty"`
	ToReleaseSessionList   []int            `json:"toReleaseSessionList,omitempty"`
	SmfChangeInfoList      []SmfChangeInfo  `json:"smfChangeInfoList,omitempty"`
	RanNodeId              *GlobalRanNodeId `json:"ranNodeId,omitempty"`
	AnN2IPv4Addr           string           `json:"anN2IPv4Addr,omitempty"`
}
