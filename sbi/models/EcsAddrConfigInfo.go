/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EcsAddrConfigInfo struct {
	Dnn                 string               `json:"dnn,omitempty"`
	Snssai              *Snssai              `json:"snssai,omitempty"`
	EcsServerAddr       *EcsServerAddr       `json:"ecsServerAddr,omitempty"`
	SpatialValidityCond *SpatialValidityCond `json:"spatialValidityCond,omitempty"`
}
