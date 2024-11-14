/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtSnssai struct {
	Sst        int         `json:"sst"`
	Sd         string      `json:"sd,omitempty"`
	SdRanges   []SdRange   `json:"sdRanges,omitempty"`
	WildcardSd *WildcardSd `json:"wildcardSd,omitempty"`
}
