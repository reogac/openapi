package models
type PolicyAssociationRequest struct {
	 HPcfId	string	`json:"hPcfId,omitempty"`
	 ServingNfId	string	`json:"servingNfId,omitempty"`
	 NotificationUri	string	`json:"notificationUri"`
	 AltNotifIpv6Addrs	[]string	`json:"altNotifIpv6Addrs,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 ServingPlmn	*PlmnIdNid	`json:"servingPlmn,omitempty"`
	 GroupIds	[]string	`json:"groupIds,omitempty"`
	 UserLoc	*UserLocation	`json:"userLoc,omitempty"`
	 TimeZone	string	`json:"timeZone,omitempty"`
	 RatType	RatType	`json:"ratType,omitempty"`
	 Guami	*Guami	`json:"guami,omitempty"`
	 ServiceName	ServiceName	`json:"serviceName,omitempty"`
	 AltNotifIpv4Addrs	[]string	`json:"altNotifIpv4Addrs,omitempty"`
	 AltNotifFqdns	[]string	`json:"altNotifFqdns,omitempty"`
	 Supi	string	`json:"supi"`
	 SuppFeat	string	`json:"suppFeat"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 AccessType	AccessType	`json:"accessType,omitempty"`
	 UePolReq	string	`json:"uePolReq,omitempty"`
	 Pc5Capab	Pc5Capability	`json:"pc5Capab,omitempty"`
	 ProSeCapab	[]string	`json:"proSeCapab,omitempty"`
}
