package models

type MmContext struct {
	UeDifferentiationInfo   *UeDifferentiationInfo `json:"ueDifferentiationInfo,omitempty"`
	TngfId                  *GlobalRanNodeId       `json:"tngfId,omitempty"`
	EpsNasSecurityMode      *EpsNasSecurityMode    `json:"epsNasSecurityMode,omitempty"`
	NasUplinkCount          *int                   `json:"nasUplinkCount,omitempty"`
	UeSecurityCapability    string                 `json:"ueSecurityCapability,omitempty"`
	NssaiMappingList        []NssaiMapping         `json:"nssaiMappingList,omitempty"`
	AllowedHomeNssai        []Snssai               `json:"allowedHomeNssai,omitempty"`
	UuaaMmStatus            string                 `json:"uuaaMmStatus,omitempty"`
	S1UeNetworkCapability   string                 `json:"s1UeNetworkCapability,omitempty"`
	ExpectedUEbehavior      *ExpectedUeBehavior    `json:"expectedUEbehavior,omitempty"`
	PlmnAssiUeRadioCapId    string                 `json:"plmnAssiUeRadioCapId,omitempty"`
	ManAssiUeRadioCapId     string                 `json:"manAssiUeRadioCapId,omitempty"`
	N3IwfId                 *GlobalRanNodeId       `json:"n3IwfId,omitempty"`
	NasSecurityMode         *NasSecurityMode       `json:"nasSecurityMode,omitempty"`
	AllowedNssai            []Snssai               `json:"allowedNssai,omitempty"`
	UcmfDicEntryId          string                 `json:"ucmfDicEntryId,omitempty"`
	AnN2ApId                *int                   `json:"anN2ApId,omitempty"`
	PendingNssaiMappingList []NssaiMapping         `json:"pendingNssaiMappingList,omitempty"`
	AccessType              AccessType             `json:"accessType"`
	NasDownlinkCount        *int                   `json:"nasDownlinkCount,omitempty"`
	NsInstanceList          []string               `json:"nsInstanceList,omitempty"`
	WagfId                  *GlobalRanNodeId       `json:"wagfId,omitempty"`
	NssaaStatusList         []NssaaStatus          `json:"nssaaStatusList,omitempty"`
}
