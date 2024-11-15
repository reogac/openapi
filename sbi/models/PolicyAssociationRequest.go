/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:26 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	TimeZone          string        `json:"timeZone,omitempty"`
	Supi              string        `json:"supi"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
}
