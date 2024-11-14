package models
type MmContext struct {
	 UuaaMmStatus	UuaaMmStatus	`json:"uuaaMmStatus,omitempty"`
	 S1UeNetworkCapability	string	`json:"s1UeNetworkCapability,omitempty"`
	 PlmnAssiUeRadioCapId	string	`json:"plmnAssiUeRadioCapId,omitempty"`
	 ManAssiUeRadioCapId	string	`json:"manAssiUeRadioCapId,omitempty"`
	 AllowedHomeNssai	[]Snssai	`json:"allowedHomeNssai,omitempty"`
	 NsInstanceList	[]string	`json:"nsInstanceList,omitempty"`
	 UeDifferentiationInfo	*UeDifferentiationInfo	`json:"ueDifferentiationInfo,omitempty"`
	 N3IwfId	*GlobalRanNodeId	`json:"n3IwfId,omitempty"`
	 WagfId	*GlobalRanNodeId	`json:"wagfId,omitempty"`
	 AccessType	AccessType	`json:"accessType"`
	 NasDownlinkCount	*int	`json:"nasDownlinkCount,omitempty"`
	 NasUplinkCount	*int	`json:"nasUplinkCount,omitempty"`
	 PendingNssaiMappingList	[]NssaiMapping	`json:"pendingNssaiMappingList,omitempty"`
	 UcmfDicEntryId	string	`json:"ucmfDicEntryId,omitempty"`
	 TngfId	*GlobalRanNodeId	`json:"tngfId,omitempty"`
	 NasSecurityMode	*NasSecurityMode	`json:"nasSecurityMode,omitempty"`
	 EpsNasSecurityMode	*EpsNasSecurityMode	`json:"epsNasSecurityMode,omitempty"`
	 AllowedNssai	[]Snssai	`json:"allowedNssai,omitempty"`
	 AnN2ApId	*int	`json:"anN2ApId,omitempty"`
	 NssaaStatusList	[]NssaaStatus	`json:"nssaaStatusList,omitempty"`
	 UeSecurityCapability	string	`json:"ueSecurityCapability,omitempty"`
	 NssaiMappingList	[]NssaiMapping	`json:"nssaiMappingList,omitempty"`
	 ExpectedUEbehavior	*ExpectedUeBehavior	`json:"expectedUEbehavior,omitempty"`
}
