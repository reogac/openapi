package models

type SmPolicyDecision struct {
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
}
