package models

type MmContext struct {
	S1UeNetworkCapability   string                 `json:"s1UeNetworkCapability,omitempty"`
	AllowedNssai            []Snssai               `json:"allowedNssai,omitempty"`
	AllowedHomeNssai        []Snssai               `json:"allowedHomeNssai,omitempty"`
	ManAssiUeRadioCapId     string                 `json:"manAssiUeRadioCapId,omitempty"`
	UcmfDicEntryId          string                 `json:"ucmfDicEntryId,omitempty"`
	TngfId                  *GlobalRanNodeId       `json:"tngfId,omitempty"`
	AccessType              string                 `json:"accessType"`
	EpsNasSecurityMode      *EpsNasSecurityMode    `json:"epsNasSecurityMode,omitempty"`
	NssaiMappingList        []NssaiMapping         `json:"nssaiMappingList,omitempty"`
	NsInstanceList          []string               `json:"nsInstanceList,omitempty"`
	ExpectedUEbehavior      *ExpectedUeBehavior    `json:"expectedUEbehavior,omitempty"`
	UeDifferentiationInfo   *UeDifferentiationInfo `json:"ueDifferentiationInfo,omitempty"`
	N3IwfId                 *GlobalRanNodeId       `json:"n3IwfId,omitempty"`
	AnN2ApId                *int                   `json:"anN2ApId,omitempty"`
	NasSecurityMode         *NasSecurityMode       `json:"nasSecurityMode,omitempty"`
	UeSecurityCapability    string                 `json:"ueSecurityCapability,omitempty"`
	PendingNssaiMappingList []NssaiMapping         `json:"pendingNssaiMappingList,omitempty"`
	UuaaMmStatus            string                 `json:"uuaaMmStatus,omitempty"`
	NasDownlinkCount        *int                   `json:"nasDownlinkCount,omitempty"`
	PlmnAssiUeRadioCapId    string                 `json:"plmnAssiUeRadioCapId,omitempty"`
	NssaaStatusList         []NssaaStatus          `json:"nssaaStatusList,omitempty"`
	NasUplinkCount          *int                   `json:"nasUplinkCount,omitempty"`
	WagfId                  *GlobalRanNodeId       `json:"wagfId,omitempty"`
}
