package models
type SmPolicyDecision struct {
	 ChargingInfo	*ChargingInformation	`json:"chargingInfo,omitempty"`
	 ReflectiveQoSTimer	*int	`json:"reflectiveQoSTimer,omitempty"`
	 RevalidationTime	string	`json:"revalidationTime,omitempty"`
	 Online	*bool	`json:"online,omitempty"`
	 QosFlowUsage	QosFlowUsage	`json:"qosFlowUsage,omitempty"`
	 PccRules	map[string]PccRule	`json:"pccRules,omitempty"`
	 QosMonDecs	map[string]QosMonitoringData	`json:"qosMonDecs,omitempty"`
	 LastReqUsageData	*RequestedUsageData	`json:"lastReqUsageData,omitempty"`
	 Ipv6Index	*int	`json:"ipv6Index,omitempty"`
	 TsnPortManContDstt	*PortManagementContainer	`json:"tsnPortManContDstt,omitempty"`
	 SessRules	map[string]SessionRule	`json:"sessRules,omitempty"`
	 QosChars	map[string]QosCharacteristics	`json:"qosChars,omitempty"`
	 PcscfRestIndication	*bool	`json:"pcscfRestIndication,omitempty"`
	 QosDecs	map[string]QosData	`json:"qosDecs,omitempty"`
	 OfflineChOnly	*bool	`json:"offlineChOnly,omitempty"`
	 LastReqRuleData	[]RequestedRuleData	`json:"lastReqRuleData,omitempty"`
	 Ipv4Index	*int	`json:"ipv4Index,omitempty"`
	 RedSessIndication	*bool	`json:"redSessIndication,omitempty"`
	 PolicyCtrlReqTriggers	[]string	`json:"policyCtrlReqTriggers,omitempty"`
	 PraInfos	map[string]PresenceInfoRm	`json:"praInfos,omitempty"`
	 UmDecs	map[string]UsageMonitoringData	`json:"umDecs,omitempty"`
	 Conds	map[string]ConditionData	`json:"conds,omitempty"`
	 RelCause	SmPolicyAssociationReleaseCause	`json:"relCause,omitempty"`
	 TsnPortManContNwtts	[]PortManagementContainer	`json:"tsnPortManContNwtts,omitempty"`
	 ChgDecs	map[string]ChargingData	`json:"chgDecs,omitempty"`
	 SuppFeat	string	`json:"suppFeat,omitempty"`
	 TsnBridgeManCont	*BridgeManagementContainer	`json:"tsnBridgeManCont,omitempty"`
	 TraffContDecs	map[string]TrafficControlData	`json:"traffContDecs,omitempty"`
	 Offline	*bool	`json:"offline,omitempty"`
}
