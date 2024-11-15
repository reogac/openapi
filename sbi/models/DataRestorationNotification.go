/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataRestorationNotification struct {
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
}
