package models

type N2InformationNotification struct {
	InitialAmfName         string           `json:"initialAmfName,omitempty"`
	AnN2IPv4Addr           string           `json:"anN2IPv4Addr,omitempty"`
	Guami                  *Guami           `json:"guami,omitempty"`
	NotifySourceNgRan      *bool            `json:"notifySourceNgRan,omitempty"`
	N2NotifySubscriptionId string           `json:"n2NotifySubscriptionId"`
	ToReleaseSessionList   []int            `json:"toReleaseSessionList,omitempty"`
	LcsCorrelationId       string           `json:"lcsCorrelationId,omitempty"`
	NotifyReason           string           `json:"notifyReason,omitempty"`
	SmfChangeInfoList      []SmfChangeInfo  `json:"smfChangeInfoList,omitempty"`
	RanNodeId              *GlobalRanNodeId `json:"ranNodeId,omitempty"`
	AnN2IPv6Addr           string           `json:"anN2IPv6Addr,omitempty"`
	N2InfoContainer        *N2InfoContainer `json:"n2InfoContainer,omitempty"`
}
