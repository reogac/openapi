/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataRestorationNotification struct {
	DnnList             []string        `json:"dnnList,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
}
