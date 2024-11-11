package models

type VsmfUpdatedData struct {
	UeTimeZone                 string                    `json:"ueTimeZone,omitempty"`
	N4InfoExt1                 *N4Information            `json:"n4InfoExt1,omitempty"`
	SecondaryRatUsageInfo      []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	QosFlowsRelList            []QosFlowItem             `json:"qosFlowsRelList,omitempty"`
	QosFlowsFailedtoRelList    []QosFlowItem             `json:"qosFlowsFailedtoRelList,omitempty"`
	N1SmInfoFromUe             *RefToBinaryData          `json:"n1SmInfoFromUe,omitempty"`
	AddUeLocation              *UserLocation             `json:"addUeLocation,omitempty"`
	FailedToAssignEbiList      []Arp                     `json:"failedToAssignEbiList,omitempty"`
	SecondaryRatUsageReport    []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	N4InfoExt2                 *N4Information            `json:"n4InfoExt2,omitempty"`
	QosFlowsAddModList         []QosFlowItem             `json:"qosFlowsAddModList,omitempty"`
	QosFlowsFailedtoAddModList []QosFlowItem             `json:"qosFlowsFailedtoAddModList,omitempty"`
	UnknownN1SmInfo            *RefToBinaryData          `json:"unknownN1SmInfo,omitempty"`
	N4Info                     *N4Information            `json:"n4Info,omitempty"`
	UeLocation                 *UserLocation             `json:"ueLocation,omitempty"`
	AssignedEbiList            []EbiArpMapping           `json:"assignedEbiList,omitempty"`
	ReleasedEbiList            []int                     `json:"releasedEbiList,omitempty"`
}
