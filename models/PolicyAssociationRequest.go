package models

type PolicyAssociationRequest struct {
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Supi              string        `json:"supi"`
	RatType           RatType       `json:"ratType,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
}
