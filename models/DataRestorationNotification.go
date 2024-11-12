package models

type DataRestorationNotification struct {
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
}
