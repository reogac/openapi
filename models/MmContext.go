package models

type MmContext struct {
	NasUplinkCount          *int                   `json:"nasUplinkCount,omitempty"`
	S1UeNetworkCapability   string                 `json:"s1UeNetworkCapability,omitempty"`
	ExpectedUEbehavior      *ExpectedUeBehavior    `json:"expectedUEbehavior,omitempty"`
	N3IwfId                 *GlobalRanNodeId       `json:"n3IwfId,omitempty"`
	AnN2ApId                *int                   `json:"anN2ApId,omitempty"`
	NasSecurityMode         *NasSecurityMode       `json:"nasSecurityMode,omitempty"`
	EpsNasSecurityMode      *EpsNasSecurityMode    `json:"epsNasSecurityMode,omitempty"`
	UeSecurityCapability    string                 `json:"ueSecurityCapability,omitempty"`
	AllowedHomeNssai        []Snssai               `json:"allowedHomeNssai,omitempty"`
	TngfId                  *GlobalRanNodeId       `json:"tngfId,omitempty"`
	NssaaStatusList         []NssaaStatus          `json:"nssaaStatusList,omitempty"`
	UuaaMmStatus            UuaaMmStatus           `json:"uuaaMmStatus,omitempty"`
	AccessType              AccessType             `json:"accessType"`
	NssaiMappingList        []NssaiMapping         `json:"nssaiMappingList,omitempty"`
	UeDifferentiationInfo   *UeDifferentiationInfo `json:"ueDifferentiationInfo,omitempty"`
	PlmnAssiUeRadioCapId    string                 `json:"plmnAssiUeRadioCapId,omitempty"`
	ManAssiUeRadioCapId     string                 `json:"manAssiUeRadioCapId,omitempty"`
	UcmfDicEntryId          string                 `json:"ucmfDicEntryId,omitempty"`
	NasDownlinkCount        *int                   `json:"nasDownlinkCount,omitempty"`
	AllowedNssai            []Snssai               `json:"allowedNssai,omitempty"`
	NsInstanceList          []string               `json:"nsInstanceList,omitempty"`
	WagfId                  *GlobalRanNodeId       `json:"wagfId,omitempty"`
	PendingNssaiMappingList []NssaiMapping         `json:"pendingNssaiMappingList,omitempty"`
}
