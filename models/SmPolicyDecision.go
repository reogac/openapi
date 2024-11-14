package models
type SmPolicyDecision struct {
	 PccRules	map[string]PccRule	`json:"pccRules,omitempty"`
	 Conds	map[string]ConditionData	`json:"conds,omitempty"`
	 RelCause	SmPolicyAssociationReleaseCause	`json:"relCause,omitempty"`
	 ChgDecs	map[string]ChargingData	`json:"chgDecs,omitempty"`
	 Offline	*bool	`json:"offline,omitempty"`
	 PraInfos	map[string]PresenceInfoRm	`json:"praInfos,omitempty"`
	 Ipv4Index	*int	`json:"ipv4Index,omitempty"`
	 RedSessIndication	*bool	`json:"redSessIndication,omitempty"`
	 UmDecs	map[string]UsageMonitoringData	`json:"umDecs,omitempty"`
	 LastReqRuleData	[]RequestedRuleData	`json:"lastReqRuleData,omitempty"`
	 TraffContDecs	map[string]TrafficControlData	`json:"traffContDecs,omitempty"`
	 ReflectiveQoSTimer	*int	`json:"reflectiveQoSTimer,omitempty"`
	 PolicyCtrlReqTriggers	[]string	`json:"policyCtrlReqTriggers,omitempty"`
	 QosDecs	map[string]QosData	`json:"qosDecs,omitempty"`
	 Online	*bool	`json:"online,omitempty"`
	 TsnPortManContDstt	*PortManagementContainer	`json:"tsnPortManContDstt,omitempty"`
	 TsnPortManContNwtts	[]PortManagementContainer	`json:"tsnPortManContNwtts,omitempty"`
	 RevalidationTime	string	`json:"revalidationTime,omitempty"`
	 OfflineChOnly	*bool	`json:"offlineChOnly,omitempty"`
	 LastReqUsageData	*RequestedUsageData	`json:"lastReqUsageData,omitempty"`
	 SessRules	map[string]SessionRule	`json:"sessRules,omitempty"`
	 PcscfRestIndication	*bool	`json:"pcscfRestIndication,omitempty"`
	 QosChars	map[string]QosCharacteristics	`json:"qosChars,omitempty"`
	 QosMonDecs	map[string]QosMonitoringData	`json:"qosMonDecs,omitempty"`
	 Ipv6Index	*int	`json:"ipv6Index,omitempty"`
	 ChargingInfo	*ChargingInformation	`json:"chargingInfo,omitempty"`
	 QosFlowUsage	QosFlowUsage	`json:"qosFlowUsage,omitempty"`
	 SuppFeat	string	`json:"suppFeat,omitempty"`
	 TsnBridgeManCont	*BridgeManagementContainer	`json:"tsnBridgeManCont,omitempty"`
}
