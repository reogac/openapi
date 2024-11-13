package models

type SmPolicyDecision struct {
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
}
