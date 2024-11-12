package models

type SmPolicyDecision struct {
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
}
