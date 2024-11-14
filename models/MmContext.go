package models
type MmContext struct {
	 UeSecurityCapability	string	`json:"ueSecurityCapability,omitempty"`
	 S1UeNetworkCapability	string	`json:"s1UeNetworkCapability,omitempty"`
	 AllowedHomeNssai	[]Snssai	`json:"allowedHomeNssai,omitempty"`
	 NsInstanceList	[]string	`json:"nsInstanceList,omitempty"`
	 ManAssiUeRadioCapId	string	`json:"manAssiUeRadioCapId,omitempty"`
	 N3IwfId	*GlobalRanNodeId	`json:"n3IwfId,omitempty"`
	 AccessType	AccessType	`json:"accessType"`
	 EpsNasSecurityMode	*EpsNasSecurityMode	`json:"epsNasSecurityMode,omitempty"`
	 AllowedNssai	[]Snssai	`json:"allowedNssai,omitempty"`
	 WagfId	*GlobalRanNodeId	`json:"wagfId,omitempty"`
	 NasSecurityMode	*NasSecurityMode	`json:"nasSecurityMode,omitempty"`
	 NasDownlinkCount	*int	`json:"nasDownlinkCount,omitempty"`
	 ExpectedUEbehavior	*ExpectedUeBehavior	`json:"expectedUEbehavior,omitempty"`
	 TngfId	*GlobalRanNodeId	`json:"tngfId,omitempty"`
	 AnN2ApId	*int	`json:"anN2ApId,omitempty"`
	 NssaaStatusList	[]NssaaStatus	`json:"nssaaStatusList,omitempty"`
	 NasUplinkCount	*int	`json:"nasUplinkCount,omitempty"`
	 NssaiMappingList	[]NssaiMapping	`json:"nssaiMappingList,omitempty"`
	 UcmfDicEntryId	string	`json:"ucmfDicEntryId,omitempty"`
	 PendingNssaiMappingList	[]NssaiMapping	`json:"pendingNssaiMappingList,omitempty"`
	 UuaaMmStatus	UuaaMmStatus	`json:"uuaaMmStatus,omitempty"`
	 UeDifferentiationInfo	*UeDifferentiationInfo	`json:"ueDifferentiationInfo,omitempty"`
	 PlmnAssiUeRadioCapId	string	`json:"plmnAssiUeRadioCapId,omitempty"`
}
