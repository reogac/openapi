package models

type N2InformationNotification struct {
	N2InfoContainer        *N2InfoContainer   `json:"n2InfoContainer,omitempty"`
	LcsCorrelationId       string             `json:"lcsCorrelationId,omitempty"`
	AnN2IPv4Addr           string             `json:"anN2IPv4Addr,omitempty"`
	AnN2IPv6Addr           string             `json:"anN2IPv6Addr,omitempty"`
	Guami                  *Guami             `json:"guami,omitempty"`
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
	NotifyReason           N2InfoNotifyReason `json:"notifyReason,omitempty"`
	SmfChangeInfoList      []SmfChangeInfo    `json:"smfChangeInfoList,omitempty"`
	RanNodeId              *GlobalRanNodeId   `json:"ranNodeId,omitempty"`
	InitialAmfName         string             `json:"initialAmfName,omitempty"`
	NotifySourceNgRan      *bool              `json:"notifySourceNgRan,omitempty"`
	ToReleaseSessionList   []int              `json:"toReleaseSessionList,omitempty"`
}
