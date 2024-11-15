/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:42 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
}
