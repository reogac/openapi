/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtSnssai struct {
	WildcardSd *WildcardSd `json:"wildcardSd,omitempty"`
	Sst        int         `json:"sst"`
	Sd         string      `json:"sd,omitempty"`
	SdRanges   []SdRange   `json:"sdRanges,omitempty"`
}
