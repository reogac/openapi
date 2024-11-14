package models
type PolicyAssociationUpdateRequest struct {
	 AltNotifIpv4Addrs	[]string	`json:"altNotifIpv4Addrs,omitempty"`
	 Guami	*Guami	`json:"guami,omitempty"`
	 ConnectState	CmState	`json:"connectState,omitempty"`
	 ServingNfId	string	`json:"servingNfId,omitempty"`
	 GroupIds	[]string	`json:"groupIds,omitempty"`
	 ProSeCapab	[]string	`json:"proSeCapab,omitempty"`
	 NotificationUri	string	`json:"notificationUri,omitempty"`
	 AltNotifIpv6Addrs	[]string	`json:"altNotifIpv6Addrs,omitempty"`
	 AltNotifFqdns	[]string	`json:"altNotifFqdns,omitempty"`
	 Triggers	[]string	`json:"triggers,omitempty"`
	 PraStatuses	map[string]PresenceInfo	`json:"praStatuses,omitempty"`
	 UePolDelResult	string	`json:"uePolDelResult,omitempty"`
	 UePolTransFailNotif	*UePolicyTransferFailureNotification	`json:"uePolTransFailNotif,omitempty"`
	 UePolReq	string	`json:"uePolReq,omitempty"`
	 PlmnId	*PlmnIdNid	`json:"plmnId,omitempty"`
	 UserLoc	*UserLocation	`json:"userLoc,omitempty"`
}
