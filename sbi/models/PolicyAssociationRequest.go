/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:42 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	RatType           RatType       `json:"ratType,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Supi              string        `json:"supi"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	GroupIds          []string      `json:"groupIds,omitempty"`
}
