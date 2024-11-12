package models

type PolicyAssociationRequest struct {
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Supi              string        `json:"supi"`
	Gpsi              string        `json:"gpsi,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	Pei               string        `json:"pei,omitempty"`
}
