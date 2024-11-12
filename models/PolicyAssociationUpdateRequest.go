package models

type PolicyAssociationUpdateRequest struct {
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
}
