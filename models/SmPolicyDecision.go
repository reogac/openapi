package models

type SmPolicyDecision struct {
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
}
