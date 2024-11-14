package models
type VsmfUpdatedData struct {
	 SecondaryRatUsageInfo	[]SecondaryRatUsageInfo	`json:"secondaryRatUsageInfo,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 ModifiedEbiListNotDelivered	*bool	`json:"modifiedEbiListNotDelivered,omitempty"`
	 N1SmInfoFromUe	*RefToBinaryData	`json:"n1SmInfoFromUe,omitempty"`
	 AssignedEbiList	[]EbiArpMapping	`json:"assignedEbiList,omitempty"`
	 SecondaryRatUsageReport	[]SecondaryRatUsageReport	`json:"secondaryRatUsageReport,omitempty"`
	 UeLocation	*UserLocation	`json:"ueLocation,omitempty"`
	 AddUeLocation	*UserLocation	`json:"addUeLocation,omitempty"`
	 FailedToAssignEbiList	[]Arp	`json:"failedToAssignEbiList,omitempty"`
	 ReleasedEbiList	[]int	`json:"releasedEbiList,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 QosFlowsAddModList	[]QosFlowItem	`json:"qosFlowsAddModList,omitempty"`
	 UnknownN1SmInfo	*RefToBinaryData	`json:"unknownN1SmInfo,omitempty"`
	 QosFlowsFailedtoRelList	[]QosFlowItem	`json:"qosFlowsFailedtoRelList,omitempty"`
	 UeTimeZone	string	`json:"ueTimeZone,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N4InfoExt3	*N4Information	`json:"n4InfoExt3,omitempty"`
	 QosFlowsRelList	[]QosFlowItem	`json:"qosFlowsRelList,omitempty"`
	 QosFlowsFailedtoAddModList	[]QosFlowItem	`json:"qosFlowsFailedtoAddModList,omitempty"`
}
