/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:13 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	TimeZone          string        `json:"timeZone,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Supi              string        `json:"supi"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	Gpsi              string        `json:"gpsi,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
}
