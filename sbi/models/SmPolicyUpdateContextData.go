/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyUpdateContextData struct {
	QosMonReports            []QosMonitoringReport        `json:"qosMonReports,omitempty"`
	InterGrpIds              []string                     `json:"interGrpIds,omitempty"`
	SatBackhaulCategory      SatelliteBackhaulCategory    `json:"satBackhaulCategory,omitempty"`
	SubsDefQos               *SubscribedDefaultQos        `json:"subsDefQos,omitempty"`
	ServNfId                 *ServingNfIdentity           `json:"servNfId,omitempty"`
	TraceReq                 *TraceData                   `json:"traceReq,omitempty"`
	AccNetChIds              []AccNetChId                 `json:"accNetChIds,omitempty"`
	AddIpv6AddrPrefixes      string                       `json:"addIpv6AddrPrefixes,omitempty"`
	UeInitResReq             *UeInitiatedResourceRequest  `json:"ueInitResReq,omitempty"`
	TsnPortManContNwtts      []PortManagementContainer    `json:"tsnPortManContNwtts,omitempty"`
	AddAccessInfo            *AdditionalAccessInfo        `json:"addAccessInfo,omitempty"`
	RelIpv6AddressPrefix     string                       `json:"relIpv6AddressPrefix,omitempty"`
	UeMac                    string                       `json:"ueMac,omitempty"`
	ThreeGppPsDataOffStatus  *bool                        `json:"3gppPsDataOffStatus,omitempty"`
	MaPduInd                 MaPduIndication              `json:"maPduInd,omitempty"`
	MulAddrInfos             []IpMulticastAddressInfo     `json:"mulAddrInfos,omitempty"`
	InvalidPolicyDecs        []InvalidParam               `json:"invalidPolicyDecs,omitempty"`
	RelUeMac                 string                       `json:"relUeMac,omitempty"`
	RelIpv4Address           string                       `json:"relIpv4Address,omitempty"`
	RefQosIndication         *bool                        `json:"refQosIndication,omitempty"`
	AddRelIpv6AddrPrefixes   string                       `json:"addRelIpv6AddrPrefixes,omitempty"`
	QncReports               []QosNotificationControlInfo `json:"qncReports,omitempty"`
	TsnBridgeInfo            *TsnBridgeInfo               `json:"tsnBridgeInfo,omitempty"`
	TsnPortManContDstt       *PortManagementContainer     `json:"tsnPortManContDstt,omitempty"`
	PccRuleId                string                       `json:"pccRuleId,omitempty"`
	ServingNetwork           *PlmnIdNid                   `json:"servingNetwork,omitempty"`
	RelAccessInfo            *AdditionalAccessInfo        `json:"relAccessInfo,omitempty"`
	UeTimeZone               string                       `json:"ueTimeZone,omitempty"`
	RuleReports              []RuleReport                 `json:"ruleReports,omitempty"`
	AccessType               AccessType                   `json:"accessType,omitempty"`
	SessRuleReports          []SessionRuleReport          `json:"sessRuleReports,omitempty"`
	NumOfPackFilter          *int                         `json:"numOfPackFilter,omitempty"`
	VplmnQosNotApp           *bool                        `json:"vplmnQosNotApp,omitempty"`
	QosFlowUsage             QosFlowUsage                 `json:"qosFlowUsage,omitempty"`
	TypesOfNotif             []string                     `json:"typesOfNotif,omitempty"`
	RatType                  RatType                      `json:"ratType,omitempty"`
	UserLocationInfoTime     string                       `json:"userLocationInfoTime,omitempty"`
	AtsssCapab               AtsssCapability              `json:"atsssCapab,omitempty"`
	PcfUeInfo                *PcfUeCallbackInfo           `json:"pcfUeInfo,omitempty"`
	Ipv6AddressPrefix        string                       `json:"ipv6AddressPrefix,omitempty"`
	AuthProfIndex            string                       `json:"authProfIndex,omitempty"`
	AppDetectionInfos        []AppDetectionInfo           `json:"appDetectionInfos,omitempty"`
	RepPraInfos              map[string]PresenceInfo      `json:"repPraInfos,omitempty"`
	RepPolicyCtrlReqTriggers []string                     `json:"repPolicyCtrlReqTriggers,omitempty"`
	TsnBridgeManCont         *BridgeManagementContainer   `json:"tsnBridgeManCont,omitempty"`
	AccuUsageReports         []AccuUsageReport            `json:"accuUsageReports,omitempty"`
	PolicyDecFailureReports  []string                     `json:"policyDecFailureReports,omitempty"`
	AnGwStatus               *bool                        `json:"anGwStatus,omitempty"`
	TrafficDescriptors       []DddTrafficDescriptor       `json:"trafficDescriptors,omitempty"`
	Ipv4Address              string                       `json:"ipv4Address,omitempty"`
	CreditManageStatus       CreditManagementStatus       `json:"creditManageStatus,omitempty"`
	NwdafDatas               []NwdafData                  `json:"nwdafDatas,omitempty"`
	UserLocationInfo         *UserLocation                `json:"userLocationInfo,omitempty"`
	SubsSessAmbr             *Ambr                        `json:"subsSessAmbr,omitempty"`
	VplmnQos                 *VplmnQos                    `json:"vplmnQos,omitempty"`
	IpDomain                 string                       `json:"ipDomain,omitempty"`
}
