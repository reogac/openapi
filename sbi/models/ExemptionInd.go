/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExemptionInd struct {
	SnssaiOnlyCongestion *bool `json:"snssaiOnlyCongestion,omitempty"`
	SnssaiDnnCongestion  *bool `json:"snssaiDnnCongestion,omitempty"`
	DnnCongestion        *bool `json:"dnnCongestion,omitempty"`
}
