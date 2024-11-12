package models

type MmContext struct {
	NssaaStatusList         []NssaaStatus          `json:"nssaaStatusList,omitempty"`
	NasDownlinkCount        *int                   `json:"nasDownlinkCount,omitempty"`
	NasUplinkCount          *int                   `json:"nasUplinkCount,omitempty"`
	ExpectedUEbehavior      *ExpectedUeBehavior    `json:"expectedUEbehavior,omitempty"`
	UcmfDicEntryId          string                 `json:"ucmfDicEntryId,omitempty"`
	PlmnAssiUeRadioCapId    string                 `json:"plmnAssiUeRadioCapId,omitempty"`
	TngfId                  *GlobalRanNodeId       `json:"tngfId,omitempty"`
	AnN2ApId                *int                   `json:"anN2ApId,omitempty"`
	NasSecurityMode         *NasSecurityMode       `json:"nasSecurityMode,omitempty"`
	AllowedNssai            []Snssai               `json:"allowedNssai,omitempty"`
	AllowedHomeNssai        []Snssai               `json:"allowedHomeNssai,omitempty"`
	UeDifferentiationInfo   *UeDifferentiationInfo `json:"ueDifferentiationInfo,omitempty"`
	N3IwfId                 *GlobalRanNodeId       `json:"n3IwfId,omitempty"`
	WagfId                  *GlobalRanNodeId       `json:"wagfId,omitempty"`
	PendingNssaiMappingList []NssaiMapping         `json:"pendingNssaiMappingList,omitempty"`
	UuaaMmStatus            UuaaMmStatus           `json:"uuaaMmStatus,omitempty"`
	EpsNasSecurityMode      *EpsNasSecurityMode    `json:"epsNasSecurityMode,omitempty"`
	UeSecurityCapability    string                 `json:"ueSecurityCapability,omitempty"`
	S1UeNetworkCapability   string                 `json:"s1UeNetworkCapability,omitempty"`
	ManAssiUeRadioCapId     string                 `json:"manAssiUeRadioCapId,omitempty"`
	AccessType              AccessType             `json:"accessType"`
	NssaiMappingList        []NssaiMapping         `json:"nssaiMappingList,omitempty"`
	NsInstanceList          []string               `json:"nsInstanceList,omitempty"`
}
