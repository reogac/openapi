/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:42 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
}
