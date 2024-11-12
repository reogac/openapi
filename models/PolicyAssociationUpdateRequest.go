package models

type PolicyAssociationUpdateRequest struct {
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
}
