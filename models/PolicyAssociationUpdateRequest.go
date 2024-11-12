package models

type PolicyAssociationUpdateRequest struct {
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
}
