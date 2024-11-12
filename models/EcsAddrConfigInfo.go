package models

type EcsAddrConfigInfo struct {
	EcsServerAddr       *EcsServerAddr       `json:"ecsServerAddr,omitempty"`
	SpatialValidityCond *SpatialValidityCond `json:"spatialValidityCond,omitempty"`
	Dnn                 string               `json:"dnn,omitempty"`
	Snssai              *Snssai              `json:"snssai,omitempty"`
}
