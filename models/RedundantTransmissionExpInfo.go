package models
type RedundantTransmissionExpInfo struct {
	 SpatialValidCon	*NetworkAreaInfo	`json:"spatialValidCon,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
	 RedTransExps	[]RedundantTransmissionExpPerTS	`json:"redTransExps"`
}
