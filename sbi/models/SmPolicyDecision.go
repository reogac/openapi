/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDecision struct {
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
}
