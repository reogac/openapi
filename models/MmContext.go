package models

type MmContext struct {
	S1UeNetworkCapability   string                 `json:"s1UeNetworkCapability,omitempty"`
	NsInstanceList          []string               `json:"nsInstanceList,omitempty"`
	PlmnAssiUeRadioCapId    string                 `json:"plmnAssiUeRadioCapId,omitempty"`
	N3IwfId                 *GlobalRanNodeId       `json:"n3IwfId,omitempty"`
	UuaaMmStatus            string                 `json:"uuaaMmStatus,omitempty"`
	ManAssiUeRadioCapId     string                 `json:"manAssiUeRadioCapId,omitempty"`
	WagfId                  *GlobalRanNodeId       `json:"wagfId,omitempty"`
	NasSecurityMode         *NasSecurityMode       `json:"nasSecurityMode,omitempty"`
	NasDownlinkCount        *int                   `json:"nasDownlinkCount,omitempty"`
	AllowedNssai            []Snssai               `json:"allowedNssai,omitempty"`
	NssaiMappingList        []NssaiMapping         `json:"nssaiMappingList,omitempty"`
	UeDifferentiationInfo   *UeDifferentiationInfo `json:"ueDifferentiationInfo,omitempty"`
	AllowedHomeNssai        []Snssai               `json:"allowedHomeNssai,omitempty"`
	TngfId                  *GlobalRanNodeId       `json:"tngfId,omitempty"`
	AnN2ApId                *int                   `json:"anN2ApId,omitempty"`
	NssaaStatusList         []NssaaStatus          `json:"nssaaStatusList,omitempty"`
	PendingNssaiMappingList []NssaiMapping         `json:"pendingNssaiMappingList,omitempty"`
	UcmfDicEntryId          string                 `json:"ucmfDicEntryId,omitempty"`
	AccessType              string                 `json:"accessType"`
	EpsNasSecurityMode      *EpsNasSecurityMode    `json:"epsNasSecurityMode,omitempty"`
	NasUplinkCount          *int                   `json:"nasUplinkCount,omitempty"`
	UeSecurityCapability    string                 `json:"ueSecurityCapability,omitempty"`
	ExpectedUEbehavior      *ExpectedUeBehavior    `json:"expectedUEbehavior,omitempty"`
}
