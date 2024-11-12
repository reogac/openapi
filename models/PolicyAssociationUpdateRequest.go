package models

type PolicyAssociationUpdateRequest struct {
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
}
