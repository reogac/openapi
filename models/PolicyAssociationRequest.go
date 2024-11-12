package models

type PolicyAssociationRequest struct {
	ServingNfId       string        `json:"servingNfId,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Supi              string        `json:"supi"`
	TimeZone          string        `json:"timeZone,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
}
