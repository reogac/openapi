/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InformationNotification struct {
	InitialAmfName         string             `json:"initialAmfName,omitempty"`
	NotifySourceNgRan      *bool              `json:"notifySourceNgRan,omitempty"`
	NotifyReason           N2InfoNotifyReason `json:"notifyReason,omitempty"`
	N2InfoContainer        *N2InfoContainer   `json:"n2InfoContainer,omitempty"`
	ToReleaseSessionList   []int              `json:"toReleaseSessionList,omitempty"`
	LcsCorrelationId       string             `json:"lcsCorrelationId,omitempty"`
	SmfChangeInfoList      []SmfChangeInfo    `json:"smfChangeInfoList,omitempty"`
	RanNodeId              *GlobalRanNodeId   `json:"ranNodeId,omitempty"`
	AnN2IPv4Addr           string             `json:"anN2IPv4Addr,omitempty"`
	AnN2IPv6Addr           string             `json:"anN2IPv6Addr,omitempty"`
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
	Guami                  *Guami             `json:"guami,omitempty"`
}
