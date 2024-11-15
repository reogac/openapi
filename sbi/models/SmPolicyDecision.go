/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDecision struct {
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
}
