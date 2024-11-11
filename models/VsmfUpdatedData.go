package models
type VsmfUpdatedData struct {
	 QosFlowsFailedtoAddModList	[]QosFlowItem	`json:"qosFlowsFailedtoAddModList,omitempty"`
	 ReleasedEbiList	[]int	`json:"releasedEbiList,omitempty"`
	 SecondaryRatUsageReport	[]SecondaryRatUsageReport	`json:"secondaryRatUsageReport,omitempty"`
	 QosFlowsFailedtoRelList	[]QosFlowItem	`json:"qosFlowsFailedtoRelList,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
	 AssignedEbiList	[]EbiArpMapping	`json:"assignedEbiList,omitempty"`
	 FailedToAssignEbiList	[]Arp	`json:"failedToAssignEbiList,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N1SmInfoFromUe	*RefToBinaryData	`json:"n1SmInfoFromUe,omitempty"`
	 UeLocation	*UserLocation	`json:"ueLocation,omitempty"`
	 AddUeLocation	*UserLocation	`json:"addUeLocation,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 QosFlowsAddModList	[]QosFlowItem	`json:"qosFlowsAddModList,omitempty"`
	 QosFlowsRelList	[]QosFlowItem	`json:"qosFlowsRelList,omitempty"`
	 UnknownN1SmInfo	*RefToBinaryData	`json:"unknownN1SmInfo,omitempty"`
	 SecondaryRatUsageInfo	[]SecondaryRatUsageInfo	`json:"secondaryRatUsageInfo,omitempty"`
}
