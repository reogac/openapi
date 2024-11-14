/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyUpdateContextData struct {
	RepPolicyCtrlReqTriggers []string                     `json:"repPolicyCtrlReqTriggers,omitempty"`
	UserLocationInfo         *UserLocation                `json:"userLocationInfo,omitempty"`
	VplmnQos                 *VplmnQos                    `json:"vplmnQos,omitempty"`
	TsnBridgeInfo            *TsnBridgeInfo               `json:"tsnBridgeInfo,omitempty"`
	SatBackhaulCategory      SatelliteBackhaulCategory    `json:"satBackhaulCategory,omitempty"`
	AccNetChIds              []AccNetChId                 `json:"accNetChIds,omitempty"`
	TsnPortManContNwtts      []PortManagementContainer    `json:"tsnPortManContNwtts,omitempty"`
	AnGwStatus               *bool                        `json:"anGwStatus,omitempty"`
	RelUeMac                 string                       `json:"relUeMac,omitempty"`
	UserLocationInfoTime     string                       `json:"userLocationInfoTime,omitempty"`
	PolicyDecFailureReports  []string                     `json:"policyDecFailureReports,omitempty"`
	InvalidPolicyDecs        []InvalidParam               `json:"invalidPolicyDecs,omitempty"`
	PccRuleId                string                       `json:"pccRuleId,omitempty"`
	TypesOfNotif             []string                     `json:"typesOfNotif,omitempty"`
	RepPraInfos              map[string]PresenceInfo      `json:"repPraInfos,omitempty"`
	ServNfId                 *ServingNfIdentity           `json:"servNfId,omitempty"`
	TsnBridgeManCont         *BridgeManagementContainer   `json:"tsnBridgeManCont,omitempty"`
	PcfUeInfo                *PcfUeCallbackInfo           `json:"pcfUeInfo,omitempty"`
	AddAccessInfo            *AdditionalAccessInfo        `json:"addAccessInfo,omitempty"`
	NwdafDatas               []NwdafData                  `json:"nwdafDatas,omitempty"`
	RatType                  RatType                      `json:"ratType,omitempty"`
	RelIpv6AddressPrefix     string                       `json:"relIpv6AddressPrefix,omitempty"`
	RuleReports              []RuleReport                 `json:"ruleReports,omitempty"`
	TrafficDescriptors       []DddTrafficDescriptor       `json:"trafficDescriptors,omitempty"`
	AddRelIpv6AddrPrefixes   string                       `json:"addRelIpv6AddrPrefixes,omitempty"`
	SubsDefQos               *SubscribedDefaultQos        `json:"subsDefQos,omitempty"`
	NumOfPackFilter          *int                         `json:"numOfPackFilter,omitempty"`
	SessRuleReports          []SessionRuleReport          `json:"sessRuleReports,omitempty"`
	QosFlowUsage             QosFlowUsage                 `json:"qosFlowUsage,omitempty"`
	AccessType               AccessType                   `json:"accessType,omitempty"`
	ServingNetwork           *PlmnIdNid                   `json:"servingNetwork,omitempty"`
	AppDetectionInfos        []AppDetectionInfo           `json:"appDetectionInfos,omitempty"`
	CreditManageStatus       CreditManagementStatus       `json:"creditManageStatus,omitempty"`
	AtsssCapab               AtsssCapability              `json:"atsssCapab,omitempty"`
	InterGrpIds              []string                     `json:"interGrpIds,omitempty"`
	IpDomain                 string                       `json:"ipDomain,omitempty"`
	UeMac                    string                       `json:"ueMac,omitempty"`
	ThreeGppPsDataOffStatus  *bool                        `json:"3gppPsDataOffStatus,omitempty"`
	MaPduInd                 MaPduIndication              `json:"maPduInd,omitempty"`
	RelIpv4Address           string                       `json:"relIpv4Address,omitempty"`
	UeTimeZone               string                       `json:"ueTimeZone,omitempty"`
	Ipv4Address              string                       `json:"ipv4Address,omitempty"`
	AccuUsageReports         []AccuUsageReport            `json:"accuUsageReports,omitempty"`
	TsnPortManContDstt       *PortManagementContainer     `json:"tsnPortManContDstt,omitempty"`
	SubsSessAmbr             *Ambr                        `json:"subsSessAmbr,omitempty"`
	AuthProfIndex            string                       `json:"authProfIndex,omitempty"`
	QncReports               []QosNotificationControlInfo `json:"qncReports,omitempty"`
	RefQosIndication         *bool                        `json:"refQosIndication,omitempty"`
	RelAccessInfo            *AdditionalAccessInfo        `json:"relAccessInfo,omitempty"`
	AddIpv6AddrPrefixes      string                       `json:"addIpv6AddrPrefixes,omitempty"`
	VplmnQosNotApp           *bool                        `json:"vplmnQosNotApp,omitempty"`
	UeInitResReq             *UeInitiatedResourceRequest  `json:"ueInitResReq,omitempty"`
	MulAddrInfos             []IpMulticastAddressInfo     `json:"mulAddrInfos,omitempty"`
	Ipv6AddressPrefix        string                       `json:"ipv6AddressPrefix,omitempty"`
	QosMonReports            []QosMonitoringReport        `json:"qosMonReports,omitempty"`
	TraceReq                 *TraceData                   `json:"traceReq,omitempty"`
}
