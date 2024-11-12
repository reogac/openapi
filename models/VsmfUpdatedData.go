package models

type VsmfUpdatedData struct {
	QosFlowsFailedtoRelList     []QosFlowItem             `json:"qosFlowsFailedtoRelList,omitempty"`
	ReleasedEbiList             []int                     `json:"releasedEbiList,omitempty"`
	N4InfoExt1                  *N4Information            `json:"n4InfoExt1,omitempty"`
	N4InfoExt2                  *N4Information            `json:"n4InfoExt2,omitempty"`
	N4InfoExt3                  *N4Information            `json:"n4InfoExt3,omitempty"`
	SecondaryRatUsageInfo       []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	N4Info                      *N4Information            `json:"n4Info,omitempty"`
	QosFlowsAddModList          []QosFlowItem             `json:"qosFlowsAddModList,omitempty"`
	QosFlowsRelList             []QosFlowItem             `json:"qosFlowsRelList,omitempty"`
	QosFlowsFailedtoAddModList  []QosFlowItem             `json:"qosFlowsFailedtoAddModList,omitempty"`
	UeTimeZone                  string                    `json:"ueTimeZone,omitempty"`
	AddUeLocation               *UserLocation             `json:"addUeLocation,omitempty"`
	SecondaryRatUsageReport     []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	ModifiedEbiListNotDelivered *bool                     `json:"modifiedEbiListNotDelivered,omitempty"`
	UnknownN1SmInfo             *RefToBinaryData          `json:"unknownN1SmInfo,omitempty"`
	UeLocation                  *UserLocation             `json:"ueLocation,omitempty"`
	AssignedEbiList             []EbiArpMapping           `json:"assignedEbiList,omitempty"`
	N1SmInfoFromUe              *RefToBinaryData          `json:"n1SmInfoFromUe,omitempty"`
	FailedToAssignEbiList       []Arp                     `json:"failedToAssignEbiList,omitempty"`
}
