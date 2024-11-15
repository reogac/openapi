/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Supi              string        `json:"supi"`
	SuppFeat          string        `json:"suppFeat"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
}
