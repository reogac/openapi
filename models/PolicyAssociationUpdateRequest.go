package models

type PolicyAssociationUpdateRequest struct {
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
}
