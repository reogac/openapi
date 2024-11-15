/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfStatusInfo struct {
	TargetAmfRemoval string       `json:"targetAmfRemoval,omitempty"`
	TargetAmfFailure string       `json:"targetAmfFailure,omitempty"`
	GuamiList        []Guami      `json:"guamiList"`
	StatusChange     StatusChange `json:"statusChange"`
}
