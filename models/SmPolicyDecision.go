package models

type SmPolicyDecision struct {
	Online                *bool                           `json:"online,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
}
