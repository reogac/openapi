/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MmContext struct {
	NasDownlinkCount        *int                   `json:"nasDownlinkCount,omitempty"`
	UeSecurityCapability    string                 `json:"ueSecurityCapability,omitempty"`
	S1UeNetworkCapability   string                 `json:"s1UeNetworkCapability,omitempty"`
	AllowedHomeNssai        []Snssai               `json:"allowedHomeNssai,omitempty"`
	NsInstanceList          []string               `json:"nsInstanceList,omitempty"`
	UeDifferentiationInfo   *UeDifferentiationInfo `json:"ueDifferentiationInfo,omitempty"`
	AccessType              AccessType             `json:"accessType"`
	NasSecurityMode         *NasSecurityMode       `json:"nasSecurityMode,omitempty"`
	WagfId                  *GlobalRanNodeId       `json:"wagfId,omitempty"`
	N3IwfId                 *GlobalRanNodeId       `json:"n3IwfId,omitempty"`
	ExpectedUEbehavior      *ExpectedUeBehavior    `json:"expectedUEbehavior,omitempty"`
	PlmnAssiUeRadioCapId    string                 `json:"plmnAssiUeRadioCapId,omitempty"`
	PendingNssaiMappingList []NssaiMapping         `json:"pendingNssaiMappingList,omitempty"`
	EpsNasSecurityMode      *EpsNasSecurityMode    `json:"epsNasSecurityMode,omitempty"`
	NssaaStatusList         []NssaaStatus          `json:"nssaaStatusList,omitempty"`
	NssaiMappingList        []NssaiMapping         `json:"nssaiMappingList,omitempty"`
	ManAssiUeRadioCapId     string                 `json:"manAssiUeRadioCapId,omitempty"`
	UcmfDicEntryId          string                 `json:"ucmfDicEntryId,omitempty"`
	TngfId                  *GlobalRanNodeId       `json:"tngfId,omitempty"`
	AnN2ApId                *int                   `json:"anN2ApId,omitempty"`
	UuaaMmStatus            UuaaMmStatus           `json:"uuaaMmStatus,omitempty"`
	NasUplinkCount          *int                   `json:"nasUplinkCount,omitempty"`
	AllowedNssai            []Snssai               `json:"allowedNssai,omitempty"`
}
