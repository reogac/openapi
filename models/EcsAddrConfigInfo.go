/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EcsAddrConfigInfo struct {
	SpatialValidityCond *SpatialValidityCond `json:"spatialValidityCond,omitempty"`
	Dnn                 string               `json:"dnn,omitempty"`
	Snssai              *Snssai              `json:"snssai,omitempty"`
	EcsServerAddr       *EcsServerAddr       `json:"ecsServerAddr,omitempty"`
}
