package models

type DataRestorationNotification struct {
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
}
