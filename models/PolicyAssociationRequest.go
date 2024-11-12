package models

type PolicyAssociationRequest struct {
	HPcfId            string        `json:"hPcfId,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Supi              string        `json:"supi"`
	RatType           RatType       `json:"ratType,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
}
