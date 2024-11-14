/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:58 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Supi              string        `json:"supi"`
	Pei               string        `json:"pei,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
}
