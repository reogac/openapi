package models

type EcsAddrConfigInfo struct {
	Snssai              *Snssai              `json:"snssai,omitempty"`
	EcsServerAddr       *EcsServerAddr       `json:"ecsServerAddr,omitempty"`
	SpatialValidityCond *SpatialValidityCond `json:"spatialValidityCond,omitempty"`
	Dnn                 string               `json:"dnn,omitempty"`
}
