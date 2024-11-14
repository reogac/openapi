package models
type DataRestorationNotification struct {
	 DnnList	[]string	`json:"dnnList,omitempty"`
	 LastReplicationTime	string	`json:"lastReplicationTime,omitempty"`
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
	 SupiRanges	[]SupiRange	`json:"supiRanges,omitempty"`
	 GpsiRanges	[]IdentityRange	`json:"gpsiRanges,omitempty"`
	 SNssaiList	[]Snssai	`json:"sNssaiList,omitempty"`
	 PlmnId	*PlmnId	`json:"plmnId,omitempty"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 UdmGroupId	string	`json:"udmGroupId,omitempty"`
}
