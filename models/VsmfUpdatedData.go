package models

type VsmfUpdatedData struct {
	N1SmInfoFromUe              *RefToBinaryData          `json:"n1SmInfoFromUe,omitempty"`
	UnknownN1SmInfo             *RefToBinaryData          `json:"unknownN1SmInfo,omitempty"`
	AddUeLocation               *UserLocation             `json:"addUeLocation,omitempty"`
	N4Info                      *N4Information            `json:"n4Info,omitempty"`
	N4InfoExt3                  *N4Information            `json:"n4InfoExt3,omitempty"`
	QosFlowsFailedtoAddModList  []QosFlowItem             `json:"qosFlowsFailedtoAddModList,omitempty"`
	QosFlowsFailedtoRelList     []QosFlowItem             `json:"qosFlowsFailedtoRelList,omitempty"`
	ReleasedEbiList             []int                     `json:"releasedEbiList,omitempty"`
	N4InfoExt1                  *N4Information            `json:"n4InfoExt1,omitempty"`
	UeLocation                  *UserLocation             `json:"ueLocation,omitempty"`
	UeTimeZone                  string                    `json:"ueTimeZone,omitempty"`
	FailedToAssignEbiList       []Arp                     `json:"failedToAssignEbiList,omitempty"`
	SecondaryRatUsageReport     []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	ModifiedEbiListNotDelivered *bool                     `json:"modifiedEbiListNotDelivered,omitempty"`
	QosFlowsAddModList          []QosFlowItem             `json:"qosFlowsAddModList,omitempty"`
	QosFlowsRelList             []QosFlowItem             `json:"qosFlowsRelList,omitempty"`
	AssignedEbiList             []EbiArpMapping           `json:"assignedEbiList,omitempty"`
	SecondaryRatUsageInfo       []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	N4InfoExt2                  *N4Information            `json:"n4InfoExt2,omitempty"`
}
