package models

type VsmfUpdatedData struct {
	UeLocation                  *UserLocation             `json:"ueLocation,omitempty"`
	FailedToAssignEbiList       []Arp                     `json:"failedToAssignEbiList,omitempty"`
	N4InfoExt2                  *N4Information            `json:"n4InfoExt2,omitempty"`
	QosFlowsFailedtoAddModList  []QosFlowItem             `json:"qosFlowsFailedtoAddModList,omitempty"`
	QosFlowsFailedtoRelList     []QosFlowItem             `json:"qosFlowsFailedtoRelList,omitempty"`
	AddUeLocation               *UserLocation             `json:"addUeLocation,omitempty"`
	AssignedEbiList             []EbiArpMapping           `json:"assignedEbiList,omitempty"`
	N4Info                      *N4Information            `json:"n4Info,omitempty"`
	N4InfoExt3                  *N4Information            `json:"n4InfoExt3,omitempty"`
	QosFlowsAddModList          []QosFlowItem             `json:"qosFlowsAddModList,omitempty"`
	QosFlowsRelList             []QosFlowItem             `json:"qosFlowsRelList,omitempty"`
	ReleasedEbiList             []int                     `json:"releasedEbiList,omitempty"`
	ModifiedEbiListNotDelivered *bool                     `json:"modifiedEbiListNotDelivered,omitempty"`
	UnknownN1SmInfo             *RefToBinaryData          `json:"unknownN1SmInfo,omitempty"`
	UeTimeZone                  string                    `json:"ueTimeZone,omitempty"`
	SecondaryRatUsageInfo       []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	N4InfoExt1                  *N4Information            `json:"n4InfoExt1,omitempty"`
	N1SmInfoFromUe              *RefToBinaryData          `json:"n1SmInfoFromUe,omitempty"`
	SecondaryRatUsageReport     []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
}
