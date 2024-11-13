package models

type DataRestorationNotification struct {
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
}
