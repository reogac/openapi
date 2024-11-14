/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InformationNotification struct {
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
	N2InfoContainer        *N2InfoContainer   `json:"n2InfoContainer,omitempty"`
	ToReleaseSessionList   []int              `json:"toReleaseSessionList,omitempty"`
	LcsCorrelationId       string             `json:"lcsCorrelationId,omitempty"`
	NotifyReason           N2InfoNotifyReason `json:"notifyReason,omitempty"`
	InitialAmfName         string             `json:"initialAmfName,omitempty"`
	AnN2IPv4Addr           string             `json:"anN2IPv4Addr,omitempty"`
	AnN2IPv6Addr           string             `json:"anN2IPv6Addr,omitempty"`
	Guami                  *Guami             `json:"guami,omitempty"`
	SmfChangeInfoList      []SmfChangeInfo    `json:"smfChangeInfoList,omitempty"`
	RanNodeId              *GlobalRanNodeId   `json:"ranNodeId,omitempty"`
	NotifySourceNgRan      *bool              `json:"notifySourceNgRan,omitempty"`
}
