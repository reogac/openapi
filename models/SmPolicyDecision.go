/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDecision struct {
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
}
