/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDecision struct {
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
}
