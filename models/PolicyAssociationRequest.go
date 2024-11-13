package models

type PolicyAssociationRequest struct {
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Supi              string        `json:"supi"`
	TimeZone          string        `json:"timeZone,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
}
