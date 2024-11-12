package models

type PolicyAssociationRequest struct {
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Supi              string        `json:"supi"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
}
