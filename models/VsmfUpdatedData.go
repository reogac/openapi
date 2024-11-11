type VsmfUpdatedData struct {
	 ReleasedEbiList	[]int	`json:"releasedEbiList,omitempty"`
	 SecondaryRatUsageReport	[]SecondaryRatUsageReport	`json:"secondaryRatUsageReport,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 QosFlowsFailedtoAddModList	[]QosFlowItem	`json:"qosFlowsFailedtoAddModList,omitempty"`
	 UnknownN1SmInfo	*RefToBinaryData	`json:"unknownN1SmInfo,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
	 AddUeLocation	*UserLocation	`json:"addUeLocation,omitempty"`
	 AssignedEbiList	[]EbiArpMapping	`json:"assignedEbiList,omitempty"`
	 QosFlowsAddModList	[]QosFlowItem	`json:"qosFlowsAddModList,omitempty"`
	 QosFlowsRelList	[]QosFlowItem	`json:"qosFlowsRelList,omitempty"`
	 UeLocation	*UserLocation	`json:"ueLocation,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 QosFlowsFailedtoRelList	[]QosFlowItem	`json:"qosFlowsFailedtoRelList,omitempty"`
	 N1SmInfoFromUe	*RefToBinaryData	`json:"n1SmInfoFromUe,omitempty"`
	 FailedToAssignEbiList	[]Arp	`json:"failedToAssignEbiList,omitempty"`
	 SecondaryRatUsageInfo	[]SecondaryRatUsageInfo	`json:"secondaryRatUsageInfo,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
}
