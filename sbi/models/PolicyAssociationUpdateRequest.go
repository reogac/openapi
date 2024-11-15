/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:13 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	Triggers            []string                             `json:"triggers,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
}
