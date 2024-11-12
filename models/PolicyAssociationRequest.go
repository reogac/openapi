package models

type PolicyAssociationRequest struct {
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Supi              string        `json:"supi"`
	Gpsi              string        `json:"gpsi,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	NotificationUri   string        `json:"notificationUri"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
}
