/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:28 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedundantTransmissionExpInfo struct {
	SpatialValidCon *NetworkAreaInfo                `json:"spatialValidCon,omitempty"`
	Dnn             string                          `json:"dnn,omitempty"`
	RedTransExps    []RedundantTransmissionExpPerTS `json:"redTransExps"`
}
