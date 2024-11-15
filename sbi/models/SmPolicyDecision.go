/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDecision struct {
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
}
