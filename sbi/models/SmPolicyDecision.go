/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:25 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDecision struct {
	Online                *bool                           `json:"online,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
}
