/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:42 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	TimeZone          string        `json:"timeZone,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	Gpsi              string        `json:"gpsi,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	Supi              string        `json:"supi"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
}
