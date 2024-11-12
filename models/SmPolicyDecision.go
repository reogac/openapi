package models

type SmPolicyDecision struct {
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
}
