/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyUpdateContextData struct {
	SubsDefQos               *SubscribedDefaultQos        `json:"subsDefQos,omitempty"`
	QosMonReports            []QosMonitoringReport        `json:"qosMonReports,omitempty"`
	TsnBridgeManCont         *BridgeManagementContainer   `json:"tsnBridgeManCont,omitempty"`
	ThreeGppPsDataOffStatus  *bool                        `json:"3gppPsDataOffStatus,omitempty"`
	AppDetectionInfos        []AppDetectionInfo           `json:"appDetectionInfos,omitempty"`
	AtsssCapab               AtsssCapability              `json:"atsssCapab,omitempty"`
	SessRuleReports          []SessionRuleReport          `json:"sessRuleReports,omitempty"`
	TsnPortManContDstt       *PortManagementContainer     `json:"tsnPortManContDstt,omitempty"`
	TypesOfNotif             []string                     `json:"typesOfNotif,omitempty"`
	Ipv6AddressPrefix        string                       `json:"ipv6AddressPrefix,omitempty"`
	UserLocationInfoTime     string                       `json:"userLocationInfoTime,omitempty"`
	SatBackhaulCategory      SatelliteBackhaulCategory    `json:"satBackhaulCategory,omitempty"`
	AddAccessInfo            *AdditionalAccessInfo        `json:"addAccessInfo,omitempty"`
	ServingNetwork           *PlmnIdNid                   `json:"servingNetwork,omitempty"`
	UeTimeZone               string                       `json:"ueTimeZone,omitempty"`
	RelIpv4Address           string                       `json:"relIpv4Address,omitempty"`
	NwdafDatas               []NwdafData                  `json:"nwdafDatas,omitempty"`
	MulAddrInfos             []IpMulticastAddressInfo     `json:"mulAddrInfos,omitempty"`
	PccRuleId                string                       `json:"pccRuleId,omitempty"`
	InterGrpIds              []string                     `json:"interGrpIds,omitempty"`
	AccessType               AccessType                   `json:"accessType,omitempty"`
	UserLocationInfo         *UserLocation                `json:"userLocationInfo,omitempty"`
	AddIpv6AddrPrefixes      string                       `json:"addIpv6AddrPrefixes,omitempty"`
	VplmnQos                 *VplmnQos                    `json:"vplmnQos,omitempty"`
	RepPraInfos              map[string]PresenceInfo      `json:"repPraInfos,omitempty"`
	QosFlowUsage             QosFlowUsage                 `json:"qosFlowUsage,omitempty"`
	AccNetChIds              []AccNetChId                 `json:"accNetChIds,omitempty"`
	AddRelIpv6AddrPrefixes   string                       `json:"addRelIpv6AddrPrefixes,omitempty"`
	AuthProfIndex            string                       `json:"authProfIndex,omitempty"`
	RatType                  RatType                      `json:"ratType,omitempty"`
	TsnBridgeInfo            *TsnBridgeInfo               `json:"tsnBridgeInfo,omitempty"`
	AnGwStatus               *bool                        `json:"anGwStatus,omitempty"`
	RepPolicyCtrlReqTriggers []string                     `json:"repPolicyCtrlReqTriggers,omitempty"`
	AccuUsageReports         []AccuUsageReport            `json:"accuUsageReports,omitempty"`
	PcfUeInfo                *PcfUeCallbackInfo           `json:"pcfUeInfo,omitempty"`
	InvalidPolicyDecs        []InvalidParam               `json:"invalidPolicyDecs,omitempty"`
	QncReports               []QosNotificationControlInfo `json:"qncReports,omitempty"`
	CreditManageStatus       CreditManagementStatus       `json:"creditManageStatus,omitempty"`
	TsnPortManContNwtts      []PortManagementContainer    `json:"tsnPortManContNwtts,omitempty"`
	RelAccessInfo            *AdditionalAccessInfo        `json:"relAccessInfo,omitempty"`
	RelUeMac                 string                       `json:"relUeMac,omitempty"`
	RuleReports              []RuleReport                 `json:"ruleReports,omitempty"`
	SubsSessAmbr             *Ambr                        `json:"subsSessAmbr,omitempty"`
	TraceReq                 *TraceData                   `json:"traceReq,omitempty"`
	Ipv4Address              string                       `json:"ipv4Address,omitempty"`
	PolicyDecFailureReports  []string                     `json:"policyDecFailureReports,omitempty"`
	VplmnQosNotApp           *bool                        `json:"vplmnQosNotApp,omitempty"`
	NumOfPackFilter          *int                         `json:"numOfPackFilter,omitempty"`
	RefQosIndication         *bool                        `json:"refQosIndication,omitempty"`
	IpDomain                 string                       `json:"ipDomain,omitempty"`
	RelIpv6AddressPrefix     string                       `json:"relIpv6AddressPrefix,omitempty"`
	UeMac                    string                       `json:"ueMac,omitempty"`
	UeInitResReq             *UeInitiatedResourceRequest  `json:"ueInitResReq,omitempty"`
	ServNfId                 *ServingNfIdentity           `json:"servNfId,omitempty"`
	MaPduInd                 MaPduIndication              `json:"maPduInd,omitempty"`
	TrafficDescriptors       []DddTrafficDescriptor       `json:"trafficDescriptors,omitempty"`
}
