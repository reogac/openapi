package models

type DataRestorationNotification struct {
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
}
