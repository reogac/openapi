package models

type SmPolicyDecision struct {
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	PccRules              map[string]PccRule              `json:"pccRules,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	ChgDecs               map[string]ChargingData         `json:"chgDecs,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	TraffContDecs         map[string]TrafficControlData   `json:"traffContDecs,omitempty"`
	UmDecs                map[string]UsageMonitoringData  `json:"umDecs,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	QosDecs               map[string]QosData              `json:"qosDecs,omitempty"`
	PraInfos              map[string]PresenceInfoRm       `json:"praInfos,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	QosChars              map[string]QosCharacteristics   `json:"qosChars,omitempty"`
	Conds                 map[string]ConditionData        `json:"conds,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	SessRules             map[string]SessionRule          `json:"sessRules,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	QosMonDecs            map[string]QosMonitoringData    `json:"qosMonDecs,omitempty"`
}
