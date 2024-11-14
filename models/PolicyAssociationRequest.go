package models
type PolicyAssociationRequest struct {
	 ServingNfId	string	`json:"servingNfId,omitempty"`
	 Pc5Capab	Pc5Capability	`json:"pc5Capab,omitempty"`
	 RatType	RatType	`json:"ratType,omitempty"`
	 GroupIds	[]string	`json:"groupIds,omitempty"`
	 ServiceName	ServiceName	`json:"serviceName,omitempty"`
	 AccessType	AccessType	`json:"accessType,omitempty"`
	 ServingPlmn	*PlmnIdNid	`json:"servingPlmn,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 UserLoc	*UserLocation	`json:"userLoc,omitempty"`
	 TimeZone	string	`json:"timeZone,omitempty"`
	 HPcfId	string	`json:"hPcfId,omitempty"`
	 UePolReq	string	`json:"uePolReq,omitempty"`
	 NotificationUri	string	`json:"notificationUri"`
	 AltNotifIpv4Addrs	[]string	`json:"altNotifIpv4Addrs,omitempty"`
	 Supi	string	`json:"supi"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 Guami	*Guami	`json:"guami,omitempty"`
	 ProSeCapab	[]string	`json:"proSeCapab,omitempty"`
	 SuppFeat	string	`json:"suppFeat"`
	 AltNotifIpv6Addrs	[]string	`json:"altNotifIpv6Addrs,omitempty"`
	 AltNotifFqdns	[]string	`json:"altNotifFqdns,omitempty"`
}
