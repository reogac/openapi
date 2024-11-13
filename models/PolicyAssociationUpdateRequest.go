package models

type PolicyAssociationUpdateRequest struct {
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
}
