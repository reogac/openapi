package models

type DataRestorationNotification struct {
	DnnList             []string        `json:"dnnList,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
}
