package models

type PolicyAssociationUpdateRequest struct {
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
}
