package models

type N2InformationNotification struct {
	SmfChangeInfoList      []SmfChangeInfo    `json:"smfChangeInfoList,omitempty"`
	RanNodeId              *GlobalRanNodeId   `json:"ranNodeId,omitempty"`
	AnN2IPv4Addr           string             `json:"anN2IPv4Addr,omitempty"`
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
	N2InfoContainer        *N2InfoContainer   `json:"n2InfoContainer,omitempty"`
	ToReleaseSessionList   []int              `json:"toReleaseSessionList,omitempty"`
	LcsCorrelationId       string             `json:"lcsCorrelationId,omitempty"`
	NotifyReason           N2InfoNotifyReason `json:"notifyReason,omitempty"`
	AnN2IPv6Addr           string             `json:"anN2IPv6Addr,omitempty"`
	Guami                  *Guami             `json:"guami,omitempty"`
	InitialAmfName         string             `json:"initialAmfName,omitempty"`
	NotifySourceNgRan      *bool              `json:"notifySourceNgRan,omitempty"`
}
