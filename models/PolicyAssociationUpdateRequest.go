package models
type PolicyAssociationUpdateRequest struct {
	 Triggers	[]string	`json:"triggers,omitempty"`
	 PlmnId	*PlmnIdNid	`json:"plmnId,omitempty"`
	 ConnectState	CmState	`json:"connectState,omitempty"`
	 UserLoc	*UserLocation	`json:"userLoc,omitempty"`
	 UePolDelResult	string	`json:"uePolDelResult,omitempty"`
	 NotificationUri	string	`json:"notificationUri,omitempty"`
	 PraStatuses	map[string]PresenceInfo	`json:"praStatuses,omitempty"`
	 UePolReq	string	`json:"uePolReq,omitempty"`
	 Guami	*Guami	`json:"guami,omitempty"`
	 GroupIds	[]string	`json:"groupIds,omitempty"`
	 ProSeCapab	[]string	`json:"proSeCapab,omitempty"`
	 AltNotifIpv4Addrs	[]string	`json:"altNotifIpv4Addrs,omitempty"`
	 AltNotifIpv6Addrs	[]string	`json:"altNotifIpv6Addrs,omitempty"`
	 AltNotifFqdns	[]string	`json:"altNotifFqdns,omitempty"`
	 UePolTransFailNotif	*UePolicyTransferFailureNotification	`json:"uePolTransFailNotif,omitempty"`
	 ServingNfId	string	`json:"servingNfId,omitempty"`
}
