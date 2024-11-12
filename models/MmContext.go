package models

type MmContext struct {
	UeSecurityCapability    string                 `json:"ueSecurityCapability,omitempty"`
	S1UeNetworkCapability   string                 `json:"s1UeNetworkCapability,omitempty"`
	AllowedHomeNssai        []Snssai               `json:"allowedHomeNssai,omitempty"`
	NsInstanceList          []string               `json:"nsInstanceList,omitempty"`
	TngfId                  *GlobalRanNodeId       `json:"tngfId,omitempty"`
	UuaaMmStatus            UuaaMmStatus           `json:"uuaaMmStatus,omitempty"`
	NasSecurityMode         *NasSecurityMode       `json:"nasSecurityMode,omitempty"`
	NasDownlinkCount        *int                   `json:"nasDownlinkCount,omitempty"`
	NasUplinkCount          *int                   `json:"nasUplinkCount,omitempty"`
	NssaiMappingList        []NssaiMapping         `json:"nssaiMappingList,omitempty"`
	AnN2ApId                *int                   `json:"anN2ApId,omitempty"`
	NssaaStatusList         []NssaaStatus          `json:"nssaaStatusList,omitempty"`
	AccessType              AccessType             `json:"accessType"`
	ExpectedUEbehavior      *ExpectedUeBehavior    `json:"expectedUEbehavior,omitempty"`
	UeDifferentiationInfo   *UeDifferentiationInfo `json:"ueDifferentiationInfo,omitempty"`
	PlmnAssiUeRadioCapId    string                 `json:"plmnAssiUeRadioCapId,omitempty"`
	ManAssiUeRadioCapId     string                 `json:"manAssiUeRadioCapId,omitempty"`
	N3IwfId                 *GlobalRanNodeId       `json:"n3IwfId,omitempty"`
	WagfId                  *GlobalRanNodeId       `json:"wagfId,omitempty"`
	EpsNasSecurityMode      *EpsNasSecurityMode    `json:"epsNasSecurityMode,omitempty"`
	UcmfDicEntryId          string                 `json:"ucmfDicEntryId,omitempty"`
	PendingNssaiMappingList []NssaiMapping         `json:"pendingNssaiMappingList,omitempty"`
	AllowedNssai            []Snssai               `json:"allowedNssai,omitempty"`
}
