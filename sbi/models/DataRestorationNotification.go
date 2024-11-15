/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataRestorationNotification struct {
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
}
