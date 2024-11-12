package models

type PolicyAssociationRequest struct {
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Supi              string        `json:"supi"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
}
