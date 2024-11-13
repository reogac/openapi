package models

type RedundantTransmissionExpInfo struct {
	Dnn             string                          `json:"dnn,omitempty"`
	RedTransExps    []RedundantTransmissionExpPerTS `json:"redTransExps"`
	SpatialValidCon *NetworkAreaInfo                `json:"spatialValidCon,omitempty"`
}
