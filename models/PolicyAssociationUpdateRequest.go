package models

type PolicyAssociationUpdateRequest struct {
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
}
