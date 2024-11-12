package models

type DataRestorationNotification struct {
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
}
