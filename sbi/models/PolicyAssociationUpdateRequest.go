/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:26 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
}
