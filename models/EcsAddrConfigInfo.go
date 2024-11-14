package models
type EcsAddrConfigInfo struct {
	 Dnn	string	`json:"dnn,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 EcsServerAddr	*EcsServerAddr	`json:"ecsServerAddr,omitempty"`
	 SpatialValidityCond	*SpatialValidityCond	`json:"spatialValidityCond,omitempty"`
}
