/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
}
