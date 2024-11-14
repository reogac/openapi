/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:58 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	Triggers            []string                             `json:"triggers,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
}
